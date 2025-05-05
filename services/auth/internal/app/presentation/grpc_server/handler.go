package grpc_server

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/application/usecases/authenticate"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/application/usecases/check_token"
	pb "github.com/VovkaGoodwin/microservices-cource/services/auth/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcHandler struct {
	pb.AuthServiceServer
	d Deps
}

type Deps struct {
	ct   check_token.UseCase
	auth authenticate.UseCase
}

func NewAuthHandler(d Deps) *GrpcHandler {
	return &GrpcHandler{
		d: d,
	}
}

func (g *GrpcHandler) CheckToken(ctx context.Context, req *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error) {
	token := req.GetToken()
	if token == "" {
		return nil, status.Error(codes.InvalidArgument, "empty token")
	}

	valid, userId, err := g.d.ct.Handle(ctx, token)
	if err != nil {
		return &pb.CheckTokenResponse{
			Valid:        valid,
			ErrorMessage: err.Error(),
		}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.CheckTokenResponse{
		Valid:  valid,
		UserId: userId,
	}, nil
}

func (g *GrpcHandler) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	if username == "" || password == "" {
		return nil, status.Error(codes.InvalidArgument, "empty credentials")
	}

	response, err := g.d.auth.Handle(ctx, authenticate.RequestDto{Login: username, Password: password})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AuthenticateResponse{Token: response.Token}, nil
}

func (g *GrpcHandler) Ping(_ context.Context, _ *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "pong",
	}, nil
}

func (g *GrpcHandler) register(server *grpc.Server) {
	pb.RegisterAuthServiceServer(server, g)
}
