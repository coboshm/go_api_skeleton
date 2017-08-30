package routing

import (
	"net/http"

	"github.com/coboshm/go_api_skeleton/pkg/handler"
	httpp "github.com/coboshm/go_api_skeleton/pkg/httpp/middleware"
)

// Route ...
type Route struct {
	Pattern     string
	Handlers    Handlers
	Middlewares httpp.Middlewares
}

// Methods ...
func (r *Route) Methods() (methods []string) {
	for _, h := range r.Handlers {
		methods = append(methods, h.Method)
	}

	return
}

// Routes is an slice of Route.
type Routes []Route

// Handler ...
type Handler struct {
	Method      string
	Handler     http.Handler
	Middlewares httpp.Middlewares
}

// Handlers ...
type Handlers []Handler

// RouteLoader loads routes in a given router.
type RouteLoader func(Router) error

// RouteFn Will run for each route.
type RouteFn func(router Router, route Route)

// DefaultLoader ...
func DefaultLoader(routes Routes) RouteLoader {
	return func(r Router) error {
		for _, route := range routes {
			registerRoute(r, route, nil)
		}

		return nil
	}
}

// WithCORSHeadersLoader adds the CORS headers to the response using all the declared methods as allowed.
// It also serves a 200 OK on the OPTIONS method.
func WithCORSHeadersLoader(routes Routes, origin string) RouteLoader {
	return func(r Router) error {
		middlewareAdder := func(r Router, route Route) {
			methods := append(route.Methods(), http.MethodOptions)
			r.Use(httpp.CORSMiddleware(origin, methods...))
		}

		for _, route := range routes {
			registerRoute(r, route, middlewareAdder)
			r.Options(route.Pattern, handler.CORSOptions(origin))
		}

		return nil
	}
}

// registerRoute registers a route in a given router with the specified pattern.
// It runs the given fn for each route.
func registerRoute(router Router, route Route, funcToApply RouteFn) {
	router.Route(route.Pattern, func(r Router) {

		// Middlewares should be registered before methods.

		r.Use(route.Middlewares...)

		if funcToApply != nil {
			funcToApply(r, route)
		}

		for _, h := range route.Handlers {
			r.Route("/", func(subRouter Router) {
				subRouter.Use(h.Middlewares...)
				subRouter.Method(h.Method, "/", h.Handler)
			})
		}
	})
}
