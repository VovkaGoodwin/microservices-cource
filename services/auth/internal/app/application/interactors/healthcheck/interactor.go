package healthcheck

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Interactor struct {
	cfg      *config.Config
	postgres *sqlx.DB
	redis    *redis.Client
}

func NewInteractor(cfg *config.Config, pdb *sqlx.DB, rdb *redis.Client) *Interactor {
	return &Interactor{
		cfg:      cfg,
		postgres: pdb,
		redis:    rdb,
	}
}
