package user

import (
	"context"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/user"
)

func (u *interactor) Ping(ctx context.Context) (string, error) {
	req := &user.PingRequest{}
	result, err := u.client.Ping(ctx, req)
	if err != nil {
		return "", err
	}

	return result.GetMessage(), nil
}
