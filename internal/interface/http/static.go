package http

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/x1m3/tertulia/internal/pkg/log"
)

// RegisterStatic add method to the mux that are not documented in the API.
func RegisterStatic(ctx context.Context, mux *chi.Mux) {
	mux.Get("/doc", documentation(ctx))
	mux.Get("/static/docs/api/api.yaml", swagger(ctx))
	mux.Get("/favicon.ico", favicon(ctx))
}

func documentation(ctx context.Context) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		writeFile(ctx, "api/spec.html", "text/html; charset=UTF-8", w)
	}
}

func favicon(ctx context.Context) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		writeFile(ctx, "api/tertulia.png", "image/png", w)
	}
}

func swagger(ctx context.Context) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		writeFile(ctx, "api/api.yaml", "text/html; charset=UTF-8", w)
	}
}

func writeFile(ctx context.Context, path string, mimeType string, w http.ResponseWriter) {
	f, err := os.ReadFile(path)
	if err != nil {
		cwd, _ := os.Getwd()
		log.Error(ctx, "cannot read file", "path", path, "error", err, "cwd", cwd)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("file not found"))
	}
	w.Header().Set("Content-Type", mimeType)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(f)
}
