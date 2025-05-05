package interactors

import (
	"context"

	"gateway/proto/auth"
)

type Auth struct {
	client auth.AuthServiceClient
}

func NewAuth(client auth.AuthServiceClient) *Auth {
	return &Auth{
		client: client,
	}
}

func (a *Auth) CheckToken(ctx context.Context, token string) (bool, string, error) {
	req := &auth.CheckTokenRequest{
		Token: token,
	}

	result, err := a.client.CheckToken(ctx, req)
	if err != nil {
		return false, "", err
	}

	return result.Valid, result.UserId, nil
}

func (a *Auth) Ping(ctx context.Context) (string, error) {
	req := &auth.PingRequest{}
	result, err := a.client.Ping(ctx, req)
	if err != nil {
		return "", err
	}

	return result.Message, err
}
