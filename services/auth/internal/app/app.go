package app

import (
	"auth/internal/config"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server *http.Server
}

func New(ctx context.Context, cfg *config.Config) *App {
	log := initLog(cfg)
	db := initPostgres(cfg, log)

	handler := initHandler(cfg, db)
	server := initHttpServer(cfg, handler)

	return &App{
		server: server,
	}
}

func (a *App) Start(ctx context.Context) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recover %s", r)
		}
	}()

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			fmt.Printf("%s", err.Error())
		}
	}()

	fmt.Print("Server started")
	<-done
	fmt.Print("Stopping server")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		fmt.Printf("server shutdown error %s", err.Error())
		return
	}

	fmt.Print("server stopped")
}
