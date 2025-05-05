package auth

import (
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/domain/contracts/clients"
	"go.uber.org/fx"
)

var Module = fx.Module("auth", fx.Provide(
	func(client clients.AuthServiceClient) *Deps {
		return &Deps{
			client: client,
		}
	},
	New,
))
