package app

import (
	"context"
	"strings"

	"github.com/google/uuid"

	pb "message/proto"
)

type GrpcHandler struct {
	pb.MessageServiceServer
}

func NewGrpcHandler() *GrpcHandler {
	return &GrpcHandler{}
}

func (g *GrpcHandler) GetMessagesByUserId(context.Context, *pb.GetMessagesByUserIdRequest) (*pb.GetMessagesByUserIdResponse, error) {
	var messages []*pb.TextItem
	for i := 0; i < 10; i++ {
		item := &pb.TextItem{
			Id:   uuid.New().String(),
			Text: strings.Repeat("message text ", i),
		}
		messages = append(messages, item)
	}

	return &pb.GetMessagesByUserIdResponse{
		Messages: messages,
	}, nil
}

func (g *GrpcHandler) SendMessage(context.Context, *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	return &pb.SendMessageResponse{
		Success: true,
	}, nil
}

func (g *GrpcHandler) Ping(context.Context, *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "pong",
	}, nil
}
