package app

import (
	"auth/internal/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	httpServer *http.Server
	rpcServer  *grpc.Server
}

func New(ctx context.Context, cfg *config.Config) *App {
	log := initLog(cfg)
	db := initPostgres(cfg, log)

	handler := initHttpHandler(cfg, db)
	httpServer := initHttpServer(cfg, handler)

	rpcHandler := initRpcHandler(cfg)
	rpcServer := initRpcServer(cfg, rpcHandler)

	return &App{
		httpServer: httpServer,
		rpcServer:  rpcServer,
	}
}

func (a *App) Start(ctx context.Context, cfg *config.Config) {
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

	slog.Info("http server started")

	go func() {
		lis, err := net.Listen(cfg.GRPCSServer.Network, fmt.Sprintf(":%s", cfg.GRPCSServer.Port))
		if err != nil {
			slog.Error("Failed to listen: %v", err)
		}
		if err := a.rpcServer.Serve(lis); err != nil {
			slog.Error("%s", err.Error())
		}
	}()
	slog.Info("info: ", slog.Any("info", a.rpcServer.GetServiceInfo()))
	slog.Info("grpc server started")
	<-done
	slog.Info("Stopping server")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("httpServer shutdown error %s", err.Error())
		return
	}

	fmt.Print("httpServer stopped")
}
