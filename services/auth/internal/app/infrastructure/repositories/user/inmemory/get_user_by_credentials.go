package inmemory

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"
	"github.com/pkg/errors"
)

func (r *Repository) GetUserByCredentials(_ context.Context, username string, password string) (models.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return *user, nil
		}
	}

	return models.User{}, errors.New("user not found")
}
