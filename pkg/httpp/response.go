package httpp

import (
	"net/http"
	"strings"
)

// AddCORSHeaders adds CORS headers.
func AddCORSHeaders(w http.ResponseWriter, origin string, methods ...string) {
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", "content-type, authorization, x-custom-auth")
}

// WriteJSONOk writes a 200 http response encoding given data.
func WriteJSONOk(w http.ResponseWriter, data interface{}) error {
	err := JSONEncoder(w, data)
	if err != nil {
		return err
	}

	return nil
}
