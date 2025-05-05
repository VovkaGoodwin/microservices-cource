package app

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/app"
	pkgHttpServer "github.com/VovkaGoodwin/microservices-cource/pkg/http_server"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/application/usecase/get_user_by_id"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/infrastructure/repository/user/inmemory"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/presentation/grpc_server"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/presentation/http_handler"
	"go.uber.org/fx"
)

func initInfrastructureLayer() fx.Option {
	return fx.Module(
		"infrastructure-layer",
		inmemory.Module,
	)
}

func initApplicationLayer() fx.Option {
	return fx.Module("application-layer",
		get_user_by_id.Module,
	)
}

func initPresentationLayer() fx.Option {
	return fx.Module("presentation-layer",
		http_handler.Module,
		grpc_server.Module,
	)
}

func initRunners() fx.Option {
	return fx.Module("runners",
		fx.Provide(func(grpc grpc_server.GrpcServer, http pkgHttpServer.HttpServer) map[string]app.Runner {
			runners := make(map[string]app.Runner, 2)
			runners["grpc"] = grpc
			runners["http"] = http

			return runners
		}),
	)
}
