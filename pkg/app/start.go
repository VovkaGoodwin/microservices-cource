package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (a *app) Start() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ctx := context.Background()

	for name, runner := range a.deps.Runners {
		runner := runner
		name := name
		go func() {
			a.deps.Log.InfoContext(ctx, "Running "+name)
			if err := runner(); err != nil {
				a.deps.Log.ErrorContext(ctx, "runner error",
					slog.String("name", name),
					slog.String("error", err.Error()),
				)
			}
		}()
	}

	<-done
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}
