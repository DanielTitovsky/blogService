package core_transport_grpc_auth_client

import (
	"context"
	"fmt"

	authgrpc "github.com/DanielTitovsky/authgrpc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type GrpcAuthClient struct {
	conn   *grpc.ClientConn
	client authgrpc.AuthGRPCClient
	config Config
}

func NewGrpcAuthClient(config Config) (*GrpcAuthClient, error) {
	conn, err := grpc.NewClient(config.addr, config.opts...)

	if err != nil {
		return nil, fmt.Errorf("Faliled to create grpc auth client: %w", err)
	}

	client := authgrpc.NewAuthGRPCClient(conn)

	return &GrpcAuthClient{
		client: client,
		conn:   conn,
		config: config,
	}, nil
}

func (gc *GrpcAuthClient) Close() error {
	return gc.conn.Close()
}

func (gc *GrpcAuthClient) ValidateToken(ctx context.Context, tokenString string) (uuid.UUID, error) {

	tokenMessage := &authgrpc.TokenString{
		Token: tokenString,
	}

	responce, err := gc.client.ValidateToken(ctx, tokenMessage)

	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid token: %w", err)
	}

	userId, err := uuid.Parse(responce.GetUserId())

	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid uuid: %w", err)
	}

	return userId, nil
}
