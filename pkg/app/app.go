package app

import "log/slog"

type App interface {
	Start()
}

var (
	_ App = (*app)(nil)
)

type (
	Runner interface {
		Start() error
	}
)

type Deps struct {
	Runners map[string]func() error
	Log     *slog.Logger
}

type app struct {
	deps *Deps
}

func newApp(deps *Deps) App {
	return &app{deps}
}
