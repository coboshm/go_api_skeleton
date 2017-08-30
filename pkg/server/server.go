package server

import (
	"fmt"

	"context"

	"net/http"

	"github.com/coboshm/go_api_skeleton/pkg/routing"
	"github.com/pkg/errors"
)

// Server serves http requests.
type Server struct {
	Config
	Router routing.Router
}

// NewServer creates new Server.
func NewServer(config Config) (*Server, error) {

	r := routing.NewChi()

	routes := routes()
	err := r.Load(routing.WithCORSHeadersLoader(routes, "*"))
	if err != nil {
		return nil, errors.Wrap(err, "Error loading the routes")
	}

	return &Server{
		Config: config,
		Router: r,
	}, nil
}

// Run ...
func (s Server) Run(ctx context.Context) error {
	return http.ListenAndServe(fmt.Sprintf(":%v", s.Port), s.Router)
}

// ServeHTTP serve just one request.
func (s Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Router.ServeHTTP(w, req)
}
