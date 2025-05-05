package token

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/services"
)

type Service struct {
	cfg *config.Config
}

func New(cfg *config.Config) services.TokenServiceInterface {
	return &Service{
		cfg: cfg,
	}
}
