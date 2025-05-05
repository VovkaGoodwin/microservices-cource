package repositories

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"
)

type AccessTokenRepositoryInterface interface {
	IsTokenRevoked(ctx context.Context, tokenId models.TokenId) (bool, error)
	SaveToken(ctx context.Context, tokenId models.TokenId, token string) error
}
