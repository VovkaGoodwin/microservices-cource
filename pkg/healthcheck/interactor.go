package healthcheck

import (
	"context"
	"github.com/jackc/pgx/v5"

	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/redis/go-redis/v9"
)

type Interactor interface {
	LivenessProbe(ctx context.Context) LivenessResponseDto
	ReadinessProbe(ctx context.Context) ReadinessResponseDto
}

var _ Interactor = (*interactor)(nil)

type deps struct {
	cfg      *config.Config
	postgres *pgx.Conn
	redis    *redis.Client
}

type interactor struct {
	*deps
}

func newInteractor(d *deps) Interactor {
	return &interactor{
		d,
	}
}
