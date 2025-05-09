package auth

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/domain/contracts/clients"
)

type Interactor interface {
	Authenticate(ctx context.Context, username string, password string) (string, error)
	Ping(ctx context.Context) (string, error)
	CheckToken(ctx context.Context, token string) (bool, string, error)
}

var _ Interactor = (*interactor)(nil)

type Deps struct {
	client clients.AuthServiceClient
}

type interactor struct {
	*Deps
}

func New(deps *Deps) Interactor {
	return &interactor{
		deps,
	}
}
