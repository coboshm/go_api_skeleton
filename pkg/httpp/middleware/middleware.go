package httpp

import "net/http"

// Middleware is an standard middleware handler with methods to compose middleware chains and http.Handler's.
type Middleware func(http.Handler) http.Handler

// Middlewares type is a slice of Middleware.
type Middlewares []func(http.Handler) http.Handler
