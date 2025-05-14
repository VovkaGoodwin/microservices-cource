package authenticate

import (
	"context"
	repoMocks "github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/repositories/mocks"
	serviceMocks "github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/services/mocks"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsecase_Handle(t *testing.T) {
	ctx := context.Background()
	user := models.NewUser("123")
	at := models.NewAccessToken(models.TokenId("token_id"), nil).WithUserId(string(user.ID()))
	userRepo := new(repoMocks.UserRepositoryInterface)
	userRepo.On("GetUserByCredentials", ctx, "test", "test").Return(*user, nil)

	tokenService := new(serviceMocks.TokenServiceInterface)
	tokenService.On("GenerateNewAccessToken", user.ID()).Return(at, nil).On("Sign", at).Return("token", nil).Once()

	accessTokenRepo := new(repoMocks.AccessTokenRepositoryInterface)
	accessTokenRepo.On("SaveToken", ctx, at.Id(), "token").Return(nil).Once()

	uc := New(Deps{
		UserRepo:        userRepo,
		AccessTokenRepo: accessTokenRepo,
		TokenService:    tokenService,
	})

	res, err := uc.Handle(ctx, RequestDto{
		Login:    "test",
		Password: "test",
	})

	assert.NoError(t, err)
	assert.Equal(t, res.Token, "token")
}
