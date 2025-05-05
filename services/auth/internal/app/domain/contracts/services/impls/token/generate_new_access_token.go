package token

import (
	"time"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

func (s *service) GenerateNewAccessToken(userId models.UserId) (*models.AccessToken, error) {
	tokenId, _ := s.generateTokenId(20)

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.AccessTokenClaims{
		UserID: string(userId),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "VovkaGoodwin",
			ID:        tokenId,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.JWT.AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	return models.NewAccessToken(models.TokenId(tokenId), unsignedToken), nil
}
