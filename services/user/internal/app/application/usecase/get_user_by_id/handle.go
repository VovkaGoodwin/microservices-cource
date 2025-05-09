package get_user_by_id

import (
	"context"
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/domain/models"
	"github.com/pkg/errors"
)

func (u *usecase) Handle(ctx context.Context, id string) (models.User, error) {
	user, err := u.userRepository.GetById(ctx, models.UserID(id))
	if err != nil {
		return models.User{}, errors.Wrap(err, "usecase")
	}

	return user, nil
}
