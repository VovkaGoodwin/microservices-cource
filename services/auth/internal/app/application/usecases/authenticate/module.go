package authenticate

import (
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/repositories"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/services"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"authenticate-usecase",
	fx.Provide(
		func(
			ur repositories.UserRepositoryInterface,
			atr repositories.AccessTokenRepositoryInterface,
			ts services.TokenServiceInterface,
		) Deps {
			return Deps{
				UserRepo:        ur,
				AccessTokenRepo: atr,
				TokenService:    ts,
			}
		},
		New,
	),
)
