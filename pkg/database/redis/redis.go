package redis

import (
	"context"
	"fmt"
	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

func newRedisDb(cfg *config.Config, lc fx.Lifecycle) *redis.Client {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DB.Redis.Host, cfg.DB.Redis.Port),
		Password: cfg.DB.Redis.Password,
		DB:       cfg.DB.Redis.Db,
	}

	rdb := redis.NewClient(options)
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(errors.Wrap(err, "redis initialization failed"))
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return rdb.Close()
		},
	})

	return rdb
}
