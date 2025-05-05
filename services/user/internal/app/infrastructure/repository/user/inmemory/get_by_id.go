package inmemory

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/domain/models"
	"github.com/pkg/errors"
)

func (r *repository) GetById(ctx context.Context, id models.UserID) (models.User, error) {
	if user, ok := r.users[id]; ok {
		return *user, nil
	}

	return models.User{}, errors.New("user not found")
}
