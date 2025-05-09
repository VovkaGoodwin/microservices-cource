package check_token

import (
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/repositories"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/services"
	"go.uber.org/fx"
)

var Module = fx.Module("check-token",
	fx.Provide(
		func(tokenRepo repositories.AccessTokenRepositoryInterface, tokenService services.TokenServiceInterface) Deps {
			return Deps{
				TokenService:    tokenService,
				TokenRepository: tokenRepo,
			}
		},
		NewUsecase,
	))
