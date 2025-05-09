package user

import (
	"context"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/domain/contracts/clients"
)

type Interactor interface {
	Ping(ctx context.Context) (string, error)
	GetUserById(ctx context.Context, userId string) (GetUserByIdResponseDto, error)
}

var _ Interactor = (*interactor)(nil)

type Deps struct {
	client clients.UserServiceClient
}

type interactor struct {
	*Deps
}

func New(deps *Deps) Interactor {
	return &interactor{
		deps,
	}
}
