package handler

import (
	"net/http"

	"github.com/coboshm/go_api_skeleton/pkg/httpp"
)

// HelloWorld returns a hello world
func HelloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := httpp.JSONEncoder(w, "HelloWorld")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
