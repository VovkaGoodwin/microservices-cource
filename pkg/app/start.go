package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func (a *app) Start(ctx context.Context) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)

	for name, runner := range a.deps.Runners {
		runner := runner
		name := name
		go func() {
			a.deps.Log.InfoContext(ctx, "Running "+name)
			if err := runner(ctx); err != nil {
				a.deps.Log.ErrorContext(ctx, "runner error",
					slog.String("name", name),
					slog.String("error", err.Error()),
				)
			}
		}()
	}

	<-done
	cancel()
}
