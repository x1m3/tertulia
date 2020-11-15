package auth

import (
	"github.com/x1m3/Tertulia/backend/internal/auth/internal/handler"
	"github.com/x1m3/Tertulia/backend/internal/server"
	"net/http"
)

func Handlers() []server.Endpoint {
	return []server.Endpoint{
		{Version: server.V1, Method: http.MethodGet, Path: "/auth_ping", Handler: handler.Ping},
	}
}
