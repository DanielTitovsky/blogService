package core_transport_http_server

import (
	"blog-service/internal/core/middleware"
	"net/http"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []middleware.Middleware
}

func NewRoute(method string, path string, handler http.HandlerFunc, middleware []middleware.Middleware) Route {
	return Route{
		Method:     method,
		Path:       path,
		Handler:    handler,
		Middleware: middleware,
	}
}
