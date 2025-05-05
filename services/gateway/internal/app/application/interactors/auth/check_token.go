package auth

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/auth"
)

func (i *interactor) CheckToken(ctx context.Context, token string) (bool, string, error) {
	req := &auth.CheckTokenRequest{
		Token: token,
	}

	result, err := i.client.CheckToken(ctx, req)
	if err != nil {
		return false, "", err
	}

	return result.Valid, result.UserId, nil
}
