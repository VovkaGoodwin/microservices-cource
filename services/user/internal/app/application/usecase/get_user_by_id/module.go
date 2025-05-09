package get_user_by_id

import (
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/domain/contracts/repositories"
	"go.uber.org/fx"
)

var Module = fx.Module("get-user-by-id",
	fx.Provide(
		func(userRepo repositories.UserRepositoryContract) *Deps {
			return &Deps{userRepository: userRepo}
		},
		New,
	),
)
