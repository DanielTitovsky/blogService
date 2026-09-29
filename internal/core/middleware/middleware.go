package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func MiddlewareChain(handler http.Handler, middleware ...Middleware) http.Handler {

	if len(middleware) != 0 {
		for i := len(middleware); i >= len(middleware); i-- {
			handler = middleware[i](handler)
		}
	}

	return handler
}
