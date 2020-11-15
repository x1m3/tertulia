package main

import (
	"context"
	"github.com/x1m3/Tertulia/backend/internal/auth"
	"github.com/x1m3/Tertulia/backend/internal/content"
	"github.com/x1m3/Tertulia/backend/internal/server"
	"github.com/x1m3/Tertulia/backend/pkg/pubsub"
	"github.com/x1m3/Tertulia/backend/pkg/service"
	"math/rand"

	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
)

const httpPort = 8080

func main() {
	rand.Seed(time.Now().UnixNano())
	logrus.SetLevel(logrus.DebugLevel)

	// Creating a context with cancellation
	ctx, cancelContext := context.WithCancel(context.Background())
	defer cancelContext()

	logrus.Info("Starting Tertulia")

	kernel := service.NewKernel(ctx, pubsub.Local(ctx), service.WithServices())
	kernel.Init()

	// Creating and starting and http server for public api
	logrus.WithFields(logrus.Fields{"port": httpPort}).Info("Tertulia started")
	apiServer := server.NewHTTPd(httpPort)

	// Registering http endpoints
	apiServer.RegisterEndpoints(content.Handlers()...)
	apiServer.RegisterEndpoints(auth.Handlers()...)

	// Starting server
	go apiServer.ListenAndServe(ctx)

	// Waiting for an OS signal cancellation
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Info("closing chronos")

	// Shutdown the servers
	apiServer.Shutdown(ctx)
	logrus.Info("chronos closed")
}
