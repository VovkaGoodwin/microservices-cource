package grpc_clients

import (
	"context"
	"log"

	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/domain/contracts/clients"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient() clients.UserServiceClient {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("user-svc.user.svc.cluster.local:50051", opts)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		return nil
	}

	client := user.NewUserServiceClient(conn)
	_, err = client.Ping(context.Background(), &user.PingRequest{})
	if err != nil {
		log.Fatalf("error pinging user service: %v", err)
		return nil
	}
	return client
}
