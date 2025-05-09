package grpc_server

import (
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	server *grpc.Server
}

func NewGrpcServer(handler *GrpcHandler) *GrpcServer {
	server := grpc.NewServer()
	handler.register(server)

	return &GrpcServer{server: server}
}

func (g *GrpcServer) Start() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	if err := g.server.Serve(lis); err != nil {
		return err
	}
	return nil
}
