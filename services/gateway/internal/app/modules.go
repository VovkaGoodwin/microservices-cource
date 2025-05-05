package app

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/app"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/application/interactors/auth"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/application/interactors/message"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/application/interactors/user"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/infrastructure/grpc_clients"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/presentation/http_server"
	"go.uber.org/fx"
)

func initInfrastructureLayer() fx.Option {
	return fx.Module(
		"infrastructure",
		grpc_clients.Module,
	)
}

func initApplicationLayer() fx.Option {
	return fx.Module(
		"application",
		auth.Module,
		message.Module,
		user.Module,
	)
}

func initPresentationLayer() fx.Option {
	return fx.Module(
		"presentation",
		http_server.Module,
	)
}

func initRuners() fx.Option {
	return fx.Module("runners",
		fx.Provide(func(http *http_server.HttpServer) map[string]app.Runner {
			runners := make(map[string]app.Runner, 1)
			runners["http"] = http

			return runners
		}),
	)
}
