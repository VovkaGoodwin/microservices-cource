package grpc_server

import (
	"context"
	"github.com/VovkaGoodwin/microservices-cource/pkg/app"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer interface {
	app.Runner
}

var _ GrpcServer = (*grpcServer)(nil)

type grpcServer struct {
	server *grpc.Server
}

func NewGrpcServer(handler *GrpcHandler) GrpcServer {
	server := grpc.NewServer()
	handler.register(server)

	return &grpcServer{server: server}
}

func (g *grpcServer) Start(_ context.Context) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	if err := g.server.Serve(lis); err != nil {
		return err
	}
	return nil
}
