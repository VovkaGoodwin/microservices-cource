package message

import (
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/domain/contracts/clients"
	"go.uber.org/fx"
)

var Module = fx.Module("message",
	fx.Provide(
		func(client clients.MessageServiceClient) *Deps {
			return &Deps{
				client: client,
			}
		},
		New,
	),
)
