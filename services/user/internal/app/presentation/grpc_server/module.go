package grpc_server

import (
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/application/interactors/healthcheck"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/application/usecase/get_user_by_id"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Module("grpc-server",
	fx.Provide(
		func(
			uc get_user_by_id.Usecase,
			hc healthcheck.Interactor,
		) *deps {
			server := grpc.NewServer()

			return &deps{
				getUserUsecase: uc,
				hc:             hc,
				server:         server,
			}
		},
		newServer,
	),
)
