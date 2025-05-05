package repositories

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"
)

type UserRepositoryInterface interface {
	GetUserByCredentials(ctx context.Context, username string, password string) (models.User, error)
}
