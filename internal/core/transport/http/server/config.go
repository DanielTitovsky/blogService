package core_transport_http_server

import "time"

type ServerMuxConfig struct {
	Addr            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

func NewServerMuxConfig() ServerMuxConfig {
	return ServerMuxConfig{
		Addr:            ":8085",
		ReadTimeout:     20 * time.Second,
		WriteTimeout:    40 * time.Second,
		ShutdownTimeout: 40 * time.Second,
	}
}
