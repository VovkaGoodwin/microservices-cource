package app

import (
	"context"
	"github.com/google/uuid"
	"message/proto"
	"strings"
)

type GrpcHandler struct {
	proto.MessageServiceServer
}

func NewGrpcHandler() *GrpcHandler {
	return &GrpcHandler{}
}

func (g *GrpcHandler) GetMessagesByUserId(context.Context, *proto.GetMessagesByUserIdRequest) (*proto.GetMessagesByUserIdResponse, error) {
	var messages []*proto.TextItem
	for i := 0; i < 10; i++ {
		item := &proto.TextItem{
			Id:   uuid.New().String(),
			Text: strings.Repeat("message text ", i),
		}
		messages = append(messages, item)
	}

	return &proto.GetMessagesByUserIdResponse{
		Messages: messages,
	}, nil
}

func (g *GrpcHandler) SendMessage(context.Context, *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	return &proto.SendMessageResponse{
		Success: true,
	}, nil
}

func (g *GrpcHandler) Ping(context.Context, *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{
		Message: "pong",
	}, nil
}
