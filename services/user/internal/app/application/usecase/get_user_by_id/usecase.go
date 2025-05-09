package get_user_by_id

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/domain/contracts/repositories"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/domain/models"
)

type Usecase interface {
	Handle(ctx context.Context, id string) (models.User, error)
}

var _ Usecase = (*usecase)(nil)

type Deps struct {
	userRepository repositories.UserRepositoryContract
}

type usecase struct {
	*Deps
}

func New(deps *Deps) Usecase {
	return &usecase{deps}
}
