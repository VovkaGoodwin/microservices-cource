package database

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDb(cfg *config.Config, log *slog.Logger) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Postgres.Host,
		cfg.DB.Postgres.Port,
		cfg.DB.Postgres.Username,
		cfg.DB.Postgres.DbName,
		cfg.DB.Postgres.Password,
		cfg.DB.Postgres.SslMode,
	)

	for i := 0; i < 10; i++ {
		db, err = sqlx.Connect("postgres", dsn)
		if err != nil {
			log.Info("database initializing failed. Retrying in 5 seconds...", slog.String("error", err.Error()))
			time.Sleep(5 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			log.Info("database initialized successfully")
			return db, nil
		}
	}

	return nil, err
}
