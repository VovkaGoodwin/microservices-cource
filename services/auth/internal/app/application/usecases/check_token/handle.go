package check_token

import (
	"context"
	"errors"
)

func (u *usecase) Handle(ctx context.Context, token string) (bool, string, error) {
	parsedToken, err := u.TokenService.ParseToken(token)
	if err != nil {
		return false, "", err
	}

	isRevoked, err := u.TokenRepository.IsTokenRevoked(ctx, parsedToken.Id())
	if err != nil {
		return false, "", err
	}

	if isRevoked {
		return false, "", errors.New("token is revoked")
	}

	return true, parsedToken.UserId, nil
}
