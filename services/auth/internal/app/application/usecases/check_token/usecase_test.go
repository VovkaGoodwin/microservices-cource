package check_token

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
	tokenService := new(serviceMocks.TokenServiceInterface)
	tokenService.On("ParseToken", "test_token").
		Return(models.NewAccessToken("id", nil).WithUserId("user_id"), nil)

	accessTokenRepo := new(repoMocks.AccessTokenRepositoryInterface)
	accessTokenRepo.On("IsTokenRevoked", ctx, models.TokenId("id")).Return(false, nil).Once()

	uc := NewUsecase(Deps{
		TokenService:    tokenService,
		TokenRepository: accessTokenRepo,
	})

	r, uid, err := uc.Handle(ctx, "test_token")
	assert.Truef(t, r, "r must be true")
	assert.Nil(t, err)
	assert.Equal(t, uid, "user_id")
}
