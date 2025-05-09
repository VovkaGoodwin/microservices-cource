package healthcheck

import (
	"context"
	"fmt"
)

func (h *Interactor) ReadinessProbe(ctx context.Context) ReadinessResponseDto {
	var pgResult int
	err := h.postgres.GetContext(ctx, &pgResult, "SELECT 1")
	if err != nil || pgResult != 1 {
		return ReadinessResponseDto{
			Result:  fmt.Sprintf("postgres check failed: %v", err),
			Success: false,
		}
	}

	redisResult := h.redis.Ping(ctx)
	if redisResult.Err() != nil {
		return ReadinessResponseDto{
			Result:  fmt.Sprintf("redis check failed: %v", err),
			Success: false,
		}
	}

	return ReadinessResponseDto{
		Result:  "ok",
		Success: true,
	}
}
