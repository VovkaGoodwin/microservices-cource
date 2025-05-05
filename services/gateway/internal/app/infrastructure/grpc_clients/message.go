package grpc_clients

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gateway/proto/message"
)

func NewMessageClient() message.MessageServiceClient {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("message-svc.message.svc.cluster.local:50051", opts)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	client := message.NewMessageServiceClient(conn)
	_, err = client.Ping(context.Background(), &message.PingRequest{})
	if err != nil {
		log.Fatalf("error connecting to message service: %v", err)
	}

	return client
}
