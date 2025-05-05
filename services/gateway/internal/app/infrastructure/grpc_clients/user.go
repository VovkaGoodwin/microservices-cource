package grpc_clients

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gateway/proto/user"
)

func NewUserClient() user.UserServiceClient {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("user-svc.user.svc.cluster.local:50051", opts)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		return nil
	}

	client := user.NewUserServiceClient(conn)
	_, err = client.Ping(context.Background(), &user.PingRequest{})
	if err != nil {
		log.Fatalf("error pinging auth service: %v", err)
		return nil
	}
	return client
}
