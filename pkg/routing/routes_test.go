package routing

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRouteKnowsItsMethods(t *testing.T) {
	r := generateSimpleRoute("/foo", http.MethodGet, http.MethodPost)

	assert.Equal(t, []string{http.MethodGet, http.MethodPost}, r.Methods())
}

func TestDefaultLoaderLoadsRoutes(t *testing.T) {
	routes := Routes{
		generateSimpleRoute("/foo", http.MethodGet, http.MethodPost),
		generateSimpleRoute("/bar", http.MethodGet, http.MethodPost),
		generateSimpleRoute("/doo", http.MethodGet, http.MethodPost),
		generateSimpleRoute("/daa", http.MethodGet, http.MethodPost),
	}

	r := NewRouterMock()
	for _, route := range routes {
		r.On("Route", route.Pattern, mock.Anything).Return(r)
	}

	l := DefaultLoader(routes)
	assert.NoError(t, l(r))
}

func TestWithCORSHeadersLoaderLoadsRoutes(t *testing.T) {
	routes := Routes{
		generateSimpleRoute("/foo", http.MethodGet, http.MethodPost),
		generateSimpleRoute("/bar", http.MethodGet, http.MethodPost),
		generateSimpleRoute("/doo", http.MethodGet, http.MethodPost),
		generateSimpleRoute("/daa", http.MethodGet, http.MethodPost),
	}

	r := NewRouterMock()
	for _, route := range routes {
		s := NewRouterMock()
		s.On("Use", mock.AnythingOfType("http.Middleware")).Once()

		r.On("Route", route.Pattern, mock.Anything).Return(s)
		r.On("Options", route.Pattern, mock.AnythingOfType("http.HandlerFunc")).Once()
	}

	l := WithCORSHeadersLoader(routes, "http://foo.bar")
	assert.NoError(t, l(r))
}

func generateSimpleRoute(pattern string, methods ...string) Route {
	handlers := Handlers{}

	for _, m := range methods {
		handlers = append(handlers, Handler{
			Method:  m,
			Handler: SuccessHandler,
		})
	}

	return Route{
		Pattern:  pattern,
		Handlers: handlers,
	}
}

// SuccessHandler Is a basic handler that returns a 200 OK code.
var SuccessHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
