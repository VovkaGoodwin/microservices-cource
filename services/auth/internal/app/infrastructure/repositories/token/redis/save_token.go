package redis

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"
)

func (r *repository) SaveToken(ctx context.Context, tokenId models.TokenId, token string) error {
	result := r.rdb.Set(ctx, string(tokenId), token, r.cfg.JWT.AccessTokenTTL)
	return result.Err()
}
