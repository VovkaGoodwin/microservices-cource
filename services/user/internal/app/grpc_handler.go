package app

import (
	"context"
	"user/proto"
)

type GrpcServer struct {
	proto.UnimplementedUserServiceServer
}

func NewGrpcHandler() *GrpcServer {
	return &GrpcServer{}
}

func (g *GrpcServer) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{
		Message: "pong",
	}, nil
}

func (g *GrpcServer) GetUserById(context.Context, *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	return &proto.GetUserResponse{
		UserId:       "123456",
		LastActivity: "25-05-2025",
		IsActive:     true,
		FirstName:    "Ivan",
		LastName:     "Ivanov",
	}, nil
}
