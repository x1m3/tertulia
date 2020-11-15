package content

import (
	"github.com/x1m3/Tertulia/backend/internal/content/internal/handler"
	"github.com/x1m3/Tertulia/backend/internal/server"
	"net/http"
)

func Handlers() []server.Endpoint {
	return []server.Endpoint{
		{Version: server.V1, Method: http.MethodGet, Path: "/ping", Handler: handler.Ping},
	}
}
