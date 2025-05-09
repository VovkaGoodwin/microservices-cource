package auth

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/auth"
)

func (i *interactor) Ping(ctx context.Context) (string, error) {
	req := &auth.PingRequest{}
	result, err := i.client.Ping(ctx, req)
	if err != nil {
		return "", err
	}

	return result.Message, err
}
