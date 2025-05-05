package postgres

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/fx"

	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func newConnect(cfg *config.Config, lc fx.Lifecycle) *pgx.Conn {
	ctx := context.Background()
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Postgres.Host,
		cfg.DB.Postgres.Port,
		cfg.DB.Postgres.Username,
		cfg.DB.Postgres.DbName,
		cfg.DB.Postgres.Password,
		cfg.DB.Postgres.SslMode,
	)

	connect, err := pgx.Connect(ctx, dsn)
	if err != nil {
		panic(errors.Wrap(err, "postgres connect error"))
	}

	err = connect.Ping(ctx)
	if err != nil {
		panic(errors.Wrap(err, "postgres ping error"))
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return connect.Close(ctx)
		},
	})

	return connect
}
