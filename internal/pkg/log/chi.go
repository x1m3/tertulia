package log

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

// HttpMiddleware installs a http middleware that logs any http request.
func HttpMiddleware(ctx context.Context) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return requestLogger(ctx)(next)
	}
}

func requestLogger(ctx context.Context) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()
			//nolint:contextcheck
			defer func() {
				ua := r.Header.Get("User-Agent")
				Info(r.Context(),
					"http req",
					"method", r.Method,
					"uri", r.RequestURI,
					"status", ww.Status(),
					"bytes", ww.BytesWritten(),
					"ua", ua,
					"d", time.Since(t1))
			}()
			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
