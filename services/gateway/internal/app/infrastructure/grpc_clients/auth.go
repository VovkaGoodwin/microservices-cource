package grpc_clients

import (
	"context"
	"log"

	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/domain/contracts/clients"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient() clients.AuthServiceClient {
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
