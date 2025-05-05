package app

import (
	"context"

	pb "user/proto"
)

type GrpcServer struct {
	pb.UnimplementedUserServiceServer
}

func NewGrpcHandler() *GrpcServer {
	return &GrpcServer{}
}

func (g *GrpcServer) Ping(_ context.Context, _ *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "pong",
	}, nil
}

func (g *GrpcServer) GetUserById(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		UserId:       "123456",
		LastActivity: "25-05-2025",
		IsActive:     true,
		FirstName:    "Ivan",
		LastName:     "Ivanov",
	}, nil
}
