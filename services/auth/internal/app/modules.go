package app

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/app"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/application/usecases/authenticate"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/application/usecases/check_token"
	tokenService "github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/services/impls/token"
	tokenRedisRepository "github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/infrastructure/repositories/token/redis"
	userInmemoryRepository "github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/infrastructure/repositories/user/inmemory"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/presentation/grpc_server"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/presentation/http_server"
	"go.uber.org/fx"
)

func initServices() fx.Option {
	return fx.Module("services",
		tokenService.Module,
	)
}

func initInfrastructureLevel() fx.Option {
	return fx.Module("infrastructure",
		tokenRedisRepository.Module,
		userInmemoryRepository.Module,
	)
}

func initPresentationLayer() fx.Option {
	return fx.Module("presentation",
		http_server.Module,
		grpc_server.Module,
	)
}

func initApplicationLayer() fx.Option {
	return fx.Module("application",
		check_token.Module,
		authenticate.Module,
	)
}

func initRunners() fx.Option {
	return fx.Module("runners",
		fx.Provide(func(http *http_server.HttpServer, grpc *grpc_server.GrpcServer) map[string]app.Runner {
			runners := make(map[string]app.Runner)
			runners["http"] = http
			runners["grpc"] = grpc

			return runners
		}),
	)
}
