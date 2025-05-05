package repositories

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/domain/models"
)

type UserRepositoryContract interface {
	GetById(ctx context.Context, id models.UserID) (models.User, error)
}
