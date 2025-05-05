package grpc_server

import (
	"context"
	pb "github.com/VovkaGoodwin/microservices-cource/services/user/proto"
)

func (s *server) Ping(ctx context.Context, _ *pb.PingRequest) (*pb.PingResponse, error) {
	result := s.hc.LivenessProbe(ctx)
	resp := &pb.PingResponse{
		Message: result.Result,
	}

	return resp, nil
}
