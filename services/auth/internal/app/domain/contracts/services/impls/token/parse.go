package token

import (
	"fmt"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

func (s *Service) ParseToken(tokenString string) (*models.AccessToken, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &models.AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.SignKey), nil
	})
	if err != nil {
		return &models.AccessToken{}, errors.Wrap(err, "invalid token")
	}

	claims, ok := parsedToken.Claims.(*models.AccessTokenClaims)
	if !ok {
		return &models.AccessToken{}, errors.New(fmt.Sprintf("invalid token claims: %+v", parsedToken.Claims))
	}

	token := models.NewAccessToken(models.TokenId(claims.ID), parsedToken).
		WithUserId(claims.UserID)

	return token, nil
}
