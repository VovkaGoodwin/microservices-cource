package app

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/VovkaGoodwin/microservices-cource/pkg/logger"
	"go.uber.org/fx"
	"log/slog"
)

func InitApp(deps ...fx.Option) *fx.App {

	opts := []fx.Option{
		fx.Provide(
			config.MustLoad,
			logger.NewLogger,
		),
	}

	opts = append(opts, deps...)
	opts = append(opts, fx.Provide(
		func(logger *slog.Logger, runners map[string]Runner) *Deps {
			r := make(map[string]func() error)
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
		app.Start()
	}))

	mainApp := fx.New(
		opts...,
	)

	return mainApp
}
