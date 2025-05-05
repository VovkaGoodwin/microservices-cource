package authenticate

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/repositories"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/services"
)

type UseCase interface {
	Handle(ctx context.Context, dto RequestDto) (ResponseDto, error)
}

var (
	_ UseCase = (*usecase)(nil)
)

type Deps struct {
	UserRepo        repositories.UserRepositoryInterface
	AccessTokenRepo repositories.AccessTokenRepositoryInterface
	TokenService    services.TokenServiceInterface
}

type usecase struct {
	deps Deps
}

func New(deps Deps) UseCase {
	return &usecase{
		deps: deps,
	}
}
