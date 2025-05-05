package healthcheck

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"healthcheck",
	fx.Provide(
		func(
			cfg *config.Config,
			postgres *pgx.Conn,
			redis *redis.Client,
		) *deps {
			return &deps{
				cfg,
				postgres,
				redis,
			}
		},
		newInteractor,
	),
)
