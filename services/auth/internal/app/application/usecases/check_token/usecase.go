package check_token

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/repositories"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/services"
)

type UseCase interface {
	Handle(ctx context.Context, token string) (bool, string, error)
}

var (
	_ UseCase = (*usecase)(nil)
)

type Deps struct {
	TokenService    services.TokenServiceInterface
	TokenRepository repositories.AccessTokenRepositoryInterface
}

type usecase struct {
	Deps
}

func NewUsecase(deps Deps) UseCase {
	return &usecase{
		Deps: deps,
	}
}
