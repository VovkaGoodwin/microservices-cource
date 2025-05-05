package token

import "github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/models"

func (s *service) Sign(token *models.AccessToken) (string, error) {
	return token.UnsignedToken().SignedString([]byte(s.cfg.JWT.SignKey))
}
