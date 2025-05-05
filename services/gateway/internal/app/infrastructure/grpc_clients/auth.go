package grpc_clients

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gateway/proto/auth"
)

func NewAuthClient() auth.AuthServiceClient {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("auth-svc.auth.svc.cluster.local:50051", opts)
	if err != nil {
		log.Fatalf("error connecting to auth service: %v", err)
	}

	client := auth.NewAuthServiceClient(conn)
	_, err = client.Ping(context.Background(), &auth.PingRequest{})
	if err != nil {
		log.Fatalf("error pinging auth service: %v", err)
	}

	return client
}
