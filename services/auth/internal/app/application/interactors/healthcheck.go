package interactors

import (
	"auth/internal/app/application/dto"
	"auth/internal/config"
	"context"
	"github.com/jmoiron/sqlx"
)

type Healthcheck struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewHealthcheck(cfg *config.Config, db *sqlx.DB) *Healthcheck {
	return &Healthcheck{
		cfg: cfg,
		db:  db,
	}
}

func (h *Healthcheck) LivenessProbe(ctx context.Context) dto.LivenessResponseDto {
	return dto.LivenessResponseDto{
		Result: "OK",
	}
}

func (h *Healthcheck) ReadinessProbe(ctx context.Context) dto.ReadinessResponseDto {
	var result int
	err := h.db.GetContext(ctx, &result, "SELECT 1")
	if err != nil || result != 1 {
		return dto.ReadinessResponseDto{
			Result: "false",
		}
	}

	return dto.ReadinessResponseDto{
		Result: "ok",
	}
}
