package core_transport_http_server

import "net/http"

type Router struct {
	Version   string
	ServerMux *http.ServeMux
}

func NewRouter(version string, mux *http.ServeMux) Router {
	return Router{
		Version:   version,
		ServerMux: mux,
	}
}

func (r *Router) RegisterRouter(routers ...Route) {

	for _, route := range routers {
		handler := http.Handler(route.Handler)

		if len(route.Middleware) != 0 {
			for i := len(route.Middleware); i >= len(route.Middleware); i-- {
				handler = route.Middleware[i](handler)
			}
		}

		fullPath := route.Method + " /api/" + r.Version + route.Path

		r.ServerMux.Handle(fullPath, handler)
	}
}
