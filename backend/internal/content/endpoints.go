package content

import (
	"github.com/x1m3/Tertulia/backend/internal/content/internal/handler"
	"github.com/x1m3/Tertulia/backend/internal/server"
	"net/http"
)

func Handlers() []server.Endpoint {
	return []server.Endpoint{
		{Version: server.V1, Method: http.MethodGet, Path: "/ping", Handler: handler.Ping},

		{Version: server.V1, Method: http.MethodGet, Path: "/posts", Handler: handler.ListPosts},
		{Version: server.V1, Method: http.MethodPost, Path: "/posts", Handler: handler.NewPost},
		{Version: server.V1, Method: http.MethodGet, Path: "/posts/:id", Handler: handler.GetPost},
		{Version: server.V1, Method: http.MethodPost, Path: "/posts/:id", Handler: handler.UpdatePost},
		{Version: server.V1, Method: http.MethodDelete, Path: "/posts/:id", Handler: handler.DeletePost},
	}
}
