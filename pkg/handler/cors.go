package handler

import (
	"net/http"

	"github.com/coboshm/go_api_skeleton/pkg/httpp"
)

// CORSOptions returns a ok response with CORS headers.
func CORSOptions(origin string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		httpp.AddCORSHeaders(w, origin, http.MethodOptions)
		w.WriteHeader(http.StatusOK)
	}
}
