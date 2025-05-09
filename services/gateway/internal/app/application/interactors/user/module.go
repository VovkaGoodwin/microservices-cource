package user

import (
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/domain/contracts/clients"
	"go.uber.org/fx"
)

var Module = fx.Module("user",
	fx.Provide(
		func(client clients.UserServiceClient) *Deps {
			return &Deps{client}
		},
		New,
	),
)
