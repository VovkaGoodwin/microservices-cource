package app

import (
	"context"
	"log/slog"

	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/VovkaGoodwin/microservices-cource/pkg/database/postgres"
	"github.com/VovkaGoodwin/microservices-cource/pkg/database/redis"
	"github.com/VovkaGoodwin/microservices-cource/pkg/healthcheck"
	"github.com/VovkaGoodwin/microservices-cource/pkg/logger"
	"go.uber.org/fx"
)

func InitApp(deps ...fx.Option) *fx.App {

	opts := []fx.Option{
		fx.Provide(
			config.MustLoad,
			logger.NewLogger,
		),
		postgres.Module,
		redis.Module,
		healthcheck.Module,
	}

	opts = append(opts, deps...)
	opts = append(opts, fx.Provide(
		func(logger *slog.Logger, runners map[string]Runner) *Deps {
			r := make(map[string]func(ctx context.Context) error)
			for name, runner := range runners {
				r[name] = runner.Start
			}

			return &Deps{
				Runners: r,
				Log:     logger,
			}
		},
	))
	opts = append(opts, fx.Provide(newApp))
	opts = append(opts, fx.Invoke(func(app App) {

	}))

	mainApp := fx.New(
		opts...,
	)

	return mainApp
}
