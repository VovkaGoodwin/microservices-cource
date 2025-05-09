package grpc_server

import (
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/application/usecases/authenticate"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/application/usecases/check_token"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"grpc",
	fx.Provide(
		func(
			ct check_token.UseCase,
			auth authenticate.UseCase,
		) Deps {
			return Deps{
				ct:   ct,
				auth: auth,
			}
		},
		NewAuthHandler,
		NewGrpcServer,
	),
)
