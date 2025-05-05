package token

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app/domain/contracts/services"
)

type service struct {
	cfg *config.Config
}

func newService(cfg *config.Config) services.TokenServiceInterface {
	return &service{
		cfg: cfg,
	}
}
