package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	api "github.com/x1m3/tertulia/internal/interface/http"
	"github.com/x1m3/tertulia/internal/pkg/config"
	"github.com/x1m3/tertulia/internal/pkg/log"
)

const shutdownTimeout = 5 * time.Second

func main() {
	ctx, cancelContext := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelContext()

	log.Config(log.LevelDebug, log.OutputJSON, os.Stdout)
	cfg, err := config.Load(ctx)
	if err != nil {
		log.Error(ctx, "cannot load config", "error", err)
		return
	}

	// ps := pubsub.NewInternalChannelPubSub()

	mux := chi.NewRouter()
	mux.Use(
		chiMiddleware.RequestID,
		log.HttpMiddleware(ctx),
		chiMiddleware.Recoverer,
		chiMiddleware.NoCache,
	)

	api.HandlerWithOptions(
		api.NewStrictHandlerWithOptions(
			api.NewServer(cfg),
			nil,
			api.StrictHTTPServerOptions{
				RequestErrorHandlerFunc:  api.RequestErrorHandlerFunc,
				ResponseErrorHandlerFunc: api.ResponseErrorHandlerFunc,
			}),
		api.ChiServerOptions{
			BaseRouter:       mux,
			ErrorHandlerFunc: api.ErrorHandlerFunc,
		})
	api.RegisterStatic(ctx, mux)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.ApiPort),
		Handler:           mux,
		ReadHeaderTimeout: cfg.RunTimeLimits.HttpReadHeaderTimeout,
		ReadTimeout:       cfg.RunTimeLimits.HttpReadTimeout,
		WriteTimeout:      cfg.RunTimeLimits.HttpWriteTimeout,
		MaxHeaderBytes:    cfg.RunTimeLimits.HttpHeaderSize,
	}

	startServers(ctx, []*http.Server{server})
	<-ctx.Done()
	shutdownServers(ctx, []*http.Server{server}, shutdownTimeout)
}

func startServers(ctx context.Context, servers []*http.Server) {
	for _, server := range servers {
		go func(ctx context.Context, server *http.Server) {
			if err := server.ListenAndServe(); err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					log.Info(ctx, "server closed", "addr", server.Addr)
					return
				}
				log.Error(ctx, "cannot start http server", "err", err)
			}
		}(ctx, server)
		log.Info(ctx, "API Listening", "addr", server.Addr)
	}
}

func shutdownServers(ctx context.Context, servers []*http.Server, timeout time.Duration) {
	shutdownContext, cancel := context.WithTimeout(ctx, timeout)
	wg := sync.WaitGroup{}
	for _, server := range servers {
		wg.Add(1)
		go func(ctx context.Context, server *http.Server) {
			defer wg.Done()
			if err := server.Shutdown(ctx); err != nil {
				log.Warn(ctx, "shutting down server", "err", err)
			}
		}(shutdownContext, server)
	}
	wg.Wait()
	cancel()
}
