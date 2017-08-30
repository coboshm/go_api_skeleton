package server

import (
	"net/http"

	"github.com/coboshm/go_api_skeleton/pkg/handler"
	"github.com/coboshm/go_api_skeleton/pkg/routing"
)

func routes() routing.Routes {
	return routing.Routes{
		{
			Pattern: "/",
			Handlers: routing.Handlers{
				{
					Method:  http.MethodGet,
					Handler: handler.HelloWorld(),
				},
			},
		},
	}
}
