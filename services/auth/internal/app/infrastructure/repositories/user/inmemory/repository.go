package inmemory

import (
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/repositories"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"
	"github.com/google/uuid"
)

type Repository struct {
	users map[models.UserId]*models.User
}

func New() repositories.UserRepositoryInterface {

	users := make(map[models.UserId]*models.User, 10)

	user := models.NewUser(models.UserId(uuid.New().String())).WithUsername("username")
	users[user.ID()] = user

	return &Repository{
		users: users,
	}
}
