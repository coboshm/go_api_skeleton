package routing

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Chi is just a wrapper of https://github.com/go-chi/chi.
// It implements our own interface (borrowed from chi).
type Chi struct {
	mux chi.Router
}

// ServeHTTP ...
func (r Chi) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

// Use ...
func (r Chi) Use(middlewares ...func(http.Handler) http.Handler) {
	r.mux.Use(middlewares...)
}

// With ...
func (r Chi) With(middlewares ...func(http.Handler) http.Handler) Router {
	return NewChiFromMux(r.mux.With(middlewares...))
}

// Group ...
func (r Chi) Group(fn func(r Router)) Router {
	if fn != nil {
		return NewChiFromMux(
			r.mux.Group(
				func(r chi.Router) {
					fn(NewChiFromMux(r))
				},
			),
		)
	}

	return NewChiFromMux(r.mux.Group(nil))
}

// Route ...
func (r Chi) Route(pattern string, fn func(r Router)) Router {
	if fn != nil {
		return NewChiFromMux(
			r.mux.Route(
				pattern,
				func(r chi.Router) {
					fn(NewChiFromMux(r))
				},
			),
		)

	}

	return NewChiFromMux(r.mux.Route(pattern, nil))
}

// Load ...
func (r Chi) Load(loader RouteLoader) error {
	return loader(r)
}

// Mount ...
func (r Chi) Mount(pattern string, h http.Handler) {
	r.mux.Mount(pattern, h)
}

// Handle ...
func (r Chi) Handle(pattern string, h http.Handler) {
	r.mux.Handle(pattern, h)
}

// HandleFunc ...
func (r Chi) HandleFunc(pattern string, h http.HandlerFunc) {
	r.mux.HandleFunc(pattern, h)
}

// Method ...
func (r Chi) Method(method, pattern string, h http.Handler) {
	r.mux.Method(method, pattern, h)
}

// MethodFunc ...
func (r Chi) MethodFunc(method, pattern string, h http.HandlerFunc) {
	r.mux.MethodFunc(method, pattern, h)
}

// Connect ...
func (r Chi) Connect(pattern string, h http.HandlerFunc) {
	r.mux.Connect(pattern, h)
}

// Delete ...
func (r Chi) Delete(pattern string, h http.HandlerFunc) {
	r.mux.Delete(pattern, h)
}

// Get ...
func (r Chi) Get(pattern string, h http.HandlerFunc) {
	r.mux.Get(pattern, h)
}

// Head ...
func (r Chi) Head(pattern string, h http.HandlerFunc) {
	r.mux.Head(pattern, h)
}

// Options ...
func (r Chi) Options(pattern string, h http.HandlerFunc) {
	r.mux.Options(pattern, h)
}

// Patch ...
func (r Chi) Patch(pattern string, h http.HandlerFunc) {
	r.mux.Patch(pattern, h)
}

// Post ...
func (r Chi) Post(pattern string, h http.HandlerFunc) {
	r.mux.Post(pattern, h)
}

// Put ...
func (r Chi) Put(pattern string, h http.HandlerFunc) {
	r.mux.Put(pattern, h)
}

// Trace ...
func (r Chi) Trace(pattern string, h http.HandlerFunc) {
	r.mux.Trace(pattern, h)
}

// NotFound ...
func (r Chi) NotFound(h http.HandlerFunc) {
	r.mux.NotFound(h)
}

// MethodNotAllowed ...
func (r Chi) MethodNotAllowed(h http.HandlerFunc) {
	r.mux.MethodNotAllowed(h)
}

// NewChi returns a Router using a chi.Router.
func NewChi() Router {
	return NewChiFromMux(chi.NewMux())
}

// NewChiFromMux returns a Router using the specified chi.Router.
func NewChiFromMux(mux chi.Router) Router {
	return Chi{mux}
}
