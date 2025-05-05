package app

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/app"
	"github.com/VovkaGoodwin/microservices-cource/pkg/http_server"
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/application/usecase/get_message_by_user_id"
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/application/usecase/ping"
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/application/usecase/send_message"
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/presentation/grpc_server"
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/presentation/http_handler"
	"go.uber.org/fx"
)

func initApplicationLayer() fx.Option {
	return fx.Module("application-layer",
		get_message_by_user_id.Module,
		send_message.Module,
		ping.Module,
	)
}

func initPresentationLayer() fx.Option {
	return fx.Module("presentation-layer",
		grpc_server.Module,
		http_handler.Module,
	)
}

func initRunners() fx.Option {
	return fx.Module("runners",
		fx.Provide(func(grpc grpc_server.GrpcServer, http http_server.HttpServer) map[string]app.Runner {
			runners := make(map[string]app.Runner)
			runners["http"] = http
			runners["grpc"] = grpc

			return runners
		}),
	)
}
