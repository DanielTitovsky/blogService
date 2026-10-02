package core_transport_grpc_auth_client

import (
	"time"

	"google.golang.org/grpc"
)

type Config struct {
	addr            string
	opts            []grpc.DialOption
	shutdownTimeOut time.Duration
}

func NewConfig(addr string, shutdownTimeOut time.Duration, opts ...grpc.DialOption) Config {
	return Config{
		addr:            addr,
		shutdownTimeOut: shutdownTimeOut,
		opts:            opts,
	}
}
