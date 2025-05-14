package get_user_by_id

import (
	"context"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/domain/contracts/repositories/mocks"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/domain/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsecase_Handle(t *testing.T) {
	ctx := context.Background()
	repo := new(mocks.UserRepositoryContract)
	repo.On("GetById", ctx, models.UserID("user_id")).Return(models.User{
		UserId:       "user_id",
		LastActivity: "01/01/2020",
		IsActive:     false,
		FirstName:    "test",
		LastName:     "test",
	}, nil)

	uc := New(&Deps{
		userRepository: repo,
	})

	user, err := uc.Handle(ctx, "user_id")
	assert.Nil(t, err)
	assert.Equal(t, models.UserID("user_id"), user.UserId)
	assert.Equal(t, "test", user.FirstName)
	assert.Equal(t, "test", user.LastName)
	assert.Equal(t, false, user.IsActive)
}
