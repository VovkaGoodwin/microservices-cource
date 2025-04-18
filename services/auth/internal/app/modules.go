package app

import (
	"auth/internal/app/application/interactors"
	"auth/internal/app/infrastructure/database"
	"auth/internal/app/presentation/http_handler"
	"auth/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envDev   = "development"
	envProd  = "production"
)

func initLog(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func initHandler(
	cfg *config.Config,
	db *sqlx.DB,
) *gin.Engine {
	handler := http_handler.NewHandler(
		cfg,
		interactors.NewHealthcheck(cfg, db),
	)

	return handler.InitRoutes()
}

func initHttpServer(cfg *config.Config, handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
		Handler:      handler,
	}

	return server
}

func initPostgres(cfg *config.Config, log *slog.Logger) *sqlx.DB {
	db, err := database.NewPostgresDb(cfg, log)
	if err != nil {
		panic("database initializing" + err.Error())
	}

	return db
}
