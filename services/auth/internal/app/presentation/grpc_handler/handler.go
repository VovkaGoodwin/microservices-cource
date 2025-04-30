package grpc_handler

import (
	pb "auth/proto"
	"context"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func (a *AuthServer) CheckToken(ctx context.Context, req *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error) {
	return &pb.CheckTokenResponse{
		Valid:        true,
		UserId:       "1234",
		ErrorMessage: "",
	}, nil
}

func (a *AuthServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "pong",
	}, nil
}
