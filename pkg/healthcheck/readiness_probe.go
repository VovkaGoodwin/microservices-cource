package healthcheck

import (
	"context"
	"fmt"
)

func (i *interactor) ReadinessProbe(ctx context.Context) ReadinessResponseDto {
	var pgResult int
	row := i.postgres.QueryRow(ctx, "SELECT 1")
	err := row.Scan(&pgResult)
	if err != nil || pgResult != 1 {
		return ReadinessResponseDto{
			Result:  fmt.Sprintf("postgres check failed: %v", err),
			Success: false,
		}
	}

	redisResult := i.redis.Ping(ctx)
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
