package grpc_server

import (
	"context"

	pb "github.com/VovkaGoodwin/microservices-cource/services/user/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.getUserUsecase.Handle(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := &pb.GetUserResponse{
		UserId:       string(user.UserId),
		LastActivity: user.LastActivity,
		LastName:     user.LastName,
		FirstName:    user.FirstName,
		IsActive:     user.IsActive,
	}
	return response, nil
}
