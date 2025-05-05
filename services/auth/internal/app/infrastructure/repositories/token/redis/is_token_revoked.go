package redis

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"
)

func (r *repository) IsTokenRevoked(ctx context.Context, tokenId models.TokenId) (bool, error) {
	result := r.rdb.Exists(ctx, string(tokenId))
	return result.Val() == 0, result.Err()
}
