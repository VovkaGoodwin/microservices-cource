package user

import (
	"context"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/user"
)

func (u *interactor) GetUserById(ctx context.Context, userId string) (GetUserByIdResponseDto, error) {
	req := &user.GetUserRequest{
		UserId: userId,
	}
	response, err := u.client.GetUserById(ctx, req)
	if err != nil {
		return GetUserByIdResponseDto{}, err
	}

	return GetUserByIdResponseDto{
		UserId:       response.UserId,
		FirstName:    response.FirstName,
		LastName:     response.LastName,
		IsActive:     response.IsActive,
		LastActivity: userId,
	}, nil
}
