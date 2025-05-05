package grpc_server

import (
	"context"
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/application/usecase/get_message_by_user_id"
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/application/usecase/ping"
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/application/usecase/send_message"
	pb "github.com/VovkaGoodwin/microservices-cource/services/message/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Deps struct {
	gmuc get_message_by_user_id.Usecase
	puc  ping.Usecase
	smuc send_message.Usecase
}

type GrpcHandler struct {
	pb.MessageServiceServer
	deps *Deps
}

func NewGrpcHandler(deps *Deps) *GrpcHandler {
	return &GrpcHandler{deps: deps}
}

func (g *GrpcHandler) GetMessagesByUserId(ctx context.Context, req *pb.GetMessagesByUserIdRequest) (*pb.GetMessagesByUserIdResponse, error) {
	r, err := g.deps.gmuc.Handle(ctx, get_message_by_user_id.GetMessageByUserIdRequest{UserId: req.GetUserId()})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var messages []*pb.TextItem
	for _, m := range r.Messages {
		messages = append(messages, &pb.TextItem{
			Text: m.Text,
			Id:   m.Id,
		})
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

func (g *GrpcHandler) Ping(ctx context.Context, _ *pb.PingRequest) (*pb.PingResponse, error) {
	r := g.deps.puc.Handle(ctx)

	return &pb.PingResponse{
		Message: r.Message,
	}, nil
}

func (g *GrpcHandler) register(server *grpc.Server) {
	pb.RegisterMessageServiceServer(server, g)
}
