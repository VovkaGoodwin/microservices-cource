package app

import (
	"context"
	"go.uber.org/fx"
	"log/slog"
)

type App interface {
	Start(ctx context.Context)
}

var (
	_ App = (*app)(nil)
)

type (
	Runner interface {
		Start(ctx context.Context) error
	}
)

type Deps struct {
	Runners map[string]func(ctx context.Context) error
	Log     *slog.Logger
}

type app struct {
	deps *Deps
}

func newApp(deps *Deps, lc fx.Lifecycle) App {
	a := &app{deps}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			a.Start(ctx)
			return nil
		},
	})
	return a
}
