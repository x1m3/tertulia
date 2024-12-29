package http

import (
	"context"

	"github.com/x1m3/tertulia/internal/pkg/config"
)

// Server is a struct that implements the API interface
type Server struct {
	cfg *config.Config
}

// HealthCheck returns info about the server health status and important dependencies
func (s Server) HealthCheck(ctx context.Context, request HealthCheckRequestObject) (HealthCheckResponseObject, error) {
	return HealthCheck200JSONResponse{Http: true}, nil
}

// NewServer is a Server constructor
func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}
