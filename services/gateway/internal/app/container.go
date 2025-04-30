package app

import (
	"gateway/internal/app/infrastructure/grpc_clients"

	auth "gateway/proto/auth"
	message "gateway/proto/message"
	user "gateway/proto/user"
)

var authClient auth.AuthServiceClient
var messageClient message.MessageServiceClient
var userClient user.UserServiceClient

func initContainer() {
	authClient = grpc_clients.NewAuthClient()
	userClient = grpc_clients.NewUserClient()
	messageClient = grpc_clients.NewMessageClient()
}

func GetAuthClient() auth.AuthServiceClient {
	return authClient
}

func GetUserClient() user.UserServiceClient {
	return userClient
}

func GetMessageClient() message.MessageServiceClient {
	return messageClient
}
