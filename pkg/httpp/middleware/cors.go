package httpp

import (
	"net/http"

	"github.com/coboshm/go_api_skeleton/pkg/httpp"
)

// CORSMiddleware adds CORS headers to the response.
func CORSMiddleware(origin string, methods ...string) Middleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			if !stringInSlice(r.Method, methods) {
				next.ServeHTTP(w, r)
				return
			}

			httpp.AddCORSHeaders(w, origin, methods...)
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
