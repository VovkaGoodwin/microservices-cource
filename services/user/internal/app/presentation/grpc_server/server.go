package grpc_server

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/app"
	"github.com/VovkaGoodwin/microservices-cource/pkg/healthcheck"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/application/usecase/get_user_by_id"
	pb "github.com/VovkaGoodwin/microservices-cource/services/user/proto"
	"google.golang.org/grpc"
)

type GrpcServer interface {
	app.Runner
	pb.UserServiceServer
}

var _ GrpcServer = (*server)(nil)

type deps struct {
	getUserUsecase get_user_by_id.Usecase
	hc             healthcheck.Interactor
	server         *grpc.Server
}

type server struct {
	pb.UnimplementedUserServiceServer
	*deps
}

func newServer(deps *deps) GrpcServer {
	s := &server{
		deps: deps,
	}

	pb.RegisterUserServiceServer(s.server, s)
	return s
}
