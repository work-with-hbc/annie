/*
Add CORS header to response.
*/

package http

import (
	"net/http"
	"strings"
)

func isPreflightRequest(r *http.Request) bool {
	return strings.ToUpper(r.Method) == "OPTIONS"
}

func handlePreflightRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	requestMethod := r.Header.Get("Access-Control-Request-Method")
	if requestMethod != "" {
		w.Header().Set("Access-Control-Allow-Method", requestMethod)
	}

	requestHeaders := r.Header.Get("Access-Control-Request-Headers")
	if requestHeaders != "" {
		w.Header().Set("Access-Control-Allow-Headers", requestHeaders)
	}
}

func handleCORSRequest(w http.ResponseWriter, r *http.Request) {
	// TODO set cors header base on request header.
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func CORSHeaderHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPreflightRequest(r) {
			handlePreflightRequest(w, r)
			return
		}

		handleCORSRequest(w, r)

		h.ServeHTTP(w, r)
	})
}
