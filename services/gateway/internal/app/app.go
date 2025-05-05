package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	httpServer *http.Server
}

func New() *App {
	initContainer()
	handler := initHttpHandler()
	server := initHttpServer(handler)

	return &App{server}
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
		if err := a.httpServer.ListenAndServe(); err != nil {
			fmt.Printf("%s", err.Error())
		}
	}()

	<-done
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	fmt.Println("http server shutdown")
}
