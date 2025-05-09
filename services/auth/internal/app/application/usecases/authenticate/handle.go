package authenticate

import (
	"context"
	"github.com/pkg/errors"
)

func (u *usecase) Handle(ctx context.Context, dto RequestDto) (ResponseDto, error) {
	user, err := u.deps.UserRepo.GetUserByCredentials(ctx, dto.Login, dto.Password)
	if err != nil {
		return ResponseDto{}, errors.Wrap(err, "cannot get user by credentials")
	}

	accessToken, err := u.deps.TokenService.GenerateNewAccessToken(user.ID())
	if err != nil {
		return ResponseDto{}, errors.Wrap(err, "cannot generate new access token")
	}

	signedToken, err := u.deps.TokenService.Sign(accessToken)
	if err != nil {
		return ResponseDto{}, errors.Wrap(err, "cannot sign new access token")
	}

	err = u.deps.AccessTokenRepo.SaveToken(ctx, accessToken.Id(), signedToken)
	if err != nil {
		return ResponseDto{}, errors.Wrap(err, "cannot save new access token")
	}

	return ResponseDto{signedToken}, nil
}
