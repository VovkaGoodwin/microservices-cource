package auth

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/auth"
)

func (i *interactor) Authenticate(ctx context.Context, username string, password string) (string, error) {
	req := &auth.AuthenticateRequest{
		Username: username,
		Password: password,
	}

	result, err := i.client.Authenticate(ctx, req)
	if err != nil {
		return "", err
	}

	return result.Token, nil
}
