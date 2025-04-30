package interactors

import (
	"context"
	"gateway/internal/app/application/dto"
	user "gateway/proto/user"
)

type User struct {
	client user.UserServiceClient
}

func NewUser(client user.UserServiceClient) *User {
	return &User{
		client: client,
	}
}

func (u *User) Ping(ctx context.Context) (string, error) {
	req := &user.PingRequest{}
	result, err := u.client.Ping(ctx, req)
	if err != nil {
		return "", err
	}

	return result.GetMessage(), nil
}

func (u *User) GetUserById(ctx context.Context, userId string) (dto.GetUserByIdResponseDto, error) {
	req := &user.GetUserRequest{
		UserId: userId,
	}
	response, err := u.client.GetUserById(ctx, req)
	if err != nil {
		return dto.GetUserByIdResponseDto{}, err
	}

	return dto.GetUserByIdResponseDto{
		UserId:       response.UserId,
		FirstName:    response.FirstName,
		LastName:     response.LastName,
		IsActive:     response.IsActive,
		LastActivity: userId,
	}, nil
}
