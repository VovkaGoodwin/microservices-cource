package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisDb(cfg *config.Config, log *slog.Logger) (*redis.Client, error) {
	var rdb *redis.Client
	var err error
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DB.Redis.Host, cfg.DB.Redis.Port),
		Password: cfg.DB.Redis.Password,
		DB:       cfg.DB.Redis.Db,
	}

	for i := 0; i < 100; i++ {
		rdb = redis.NewClient(options)

		result := rdb.Ping(context.Background())
		err = result.Err()
		if err != nil {
			log.Info("redis initialized failed. Retrying in 5 seconds...", slog.String("error", err.Error()))
			time.Sleep(5 * time.Second)
			continue
		}

		log.Info("redis initialized successfully")
		return rdb, nil
	}

	return nil, err
}
