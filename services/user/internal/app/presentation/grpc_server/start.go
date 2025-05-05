package grpc_server

import (
	"context"
	"github.com/pkg/errors"
	"net"
)

func (s *server) Start(_ context.Context) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	if err := s.server.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}
	return nil
}
