package services

import "github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"

type TokenServiceInterface interface {
	ParseToken(tokenString string) (*models.AccessToken, error)
	Sign(token *models.AccessToken) (string, error)
	GenerateNewAccessToken(userId models.UserId) (*models.AccessToken, error)
}
