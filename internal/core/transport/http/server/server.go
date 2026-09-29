package core_transport_http_server

import (
	"blog-service/internal/core/middleware"
	"context"
	"errors"
	"fmt"
	"net/http"
)

type ServerMux struct {
	config     ServerMuxConfig
	ServerMux  *http.ServeMux
	Middleware []middleware.Middleware
}

func NewServerMux(config ServerMuxConfig, middleware ...middleware.Middleware) *ServerMux {
	return &ServerMux{
		config:     config,
		ServerMux:  http.NewServeMux(),
		Middleware: middleware,
	}
}

func (s *ServerMux) Run(ctx context.Context) error {
	mux := middleware.MiddlewareChain(s.ServerMux, s.Middleware...)

	server := http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		ch <- server.ListenAndServe()
	}()

	select {
	case err := <-ch:
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("listen and Serve: %w", err)
		}
	case <-ctx.Done():
		fmt.Println("jopa: shutdown initiated")
		shutgownContext, cancel := context.WithTimeout(
			context.Background(),
			s.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutgownContext); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown http server %w", err)
		}

		fmt.Println("jopa: shutdown completed successfully")
	}

	return nil
}
