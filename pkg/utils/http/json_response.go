/*
JSON responses
*/

package http

import (
	"fmt"
	"net/http"
)

func makeJsonResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func JsonResponseHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		makeJsonResponse(w)
		h.ServeHTTP(w, r)
	})
}

func makeErrorResponse(w http.ResponseWriter, status int, reason string) {
	// Ensure this is a json response.
	makeJsonResponse(w)
	w.WriteHeader(status)
	fmt.Fprintf(
		w,
		"{\"reason\": \"%s\"}",
		reason,
	)
}

func InvalidInput(w http.ResponseWriter) {
	makeErrorResponse(w, 400, "invalid input")
}

func NotFound(w http.ResponseWriter) {
	makeErrorResponse(w, 404, "not found")
}

func ServerError(w http.ResponseWriter) {
	makeErrorResponse(w, 500, "server error")
}
