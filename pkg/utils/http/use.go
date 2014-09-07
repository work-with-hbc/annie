/*
Use a http middleware.
*/

package http

import (
	"net/http"
)

type MiddlewareFunc func(http.Handler) http.Handler

func HandlerFuncUse(base http.HandlerFunc, middlewares ...MiddlewareFunc) http.Handler {
	var handler http.Handler = base
	return HandlerUse(handler, middlewares...)
}

func HandlerUse(handler http.Handler, middlewares ...MiddlewareFunc) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
