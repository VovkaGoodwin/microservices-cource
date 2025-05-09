package inmemory

import (
	"fmt"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/domain/contracts/repositories"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/domain/models"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

var _ repositories.UserRepositoryContract = (*repository)(nil)

type repository struct {
	users map[models.UserID]*models.User
}

func New() repositories.UserRepositoryContract {

	users := make(map[models.UserID]*models.User, 10)

	for i := 0; i < 10; i++ {
		user := &models.User{
			UserId:       models.UserID(uuid.New().String()),
			LastName:     fmt.Sprintf("User last name %d", i+1),
			FirstName:    fmt.Sprintf("User first name %d", i+1),
			IsActive:     rand.Intn(2)%2 == 0,
			LastActivity: time.Now().GoString(),
		}
		users[user.UserId] = user
	}

	return &repository{
		users: users,
	}
}
