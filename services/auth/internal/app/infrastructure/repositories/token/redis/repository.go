package redis

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/repositories"
	"github.com/redis/go-redis/v9"
)

type repository struct {
	rdb *redis.Client
	cfg *config.Config
}

func newRepository(cfg *config.Config, redis *redis.Client) repositories.AccessTokenRepositoryInterface {
	return &repository{
		rdb: redis,
		cfg: cfg,
	}
}
