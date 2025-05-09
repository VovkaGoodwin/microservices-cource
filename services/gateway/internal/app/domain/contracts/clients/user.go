package clients

import "github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/user"

type UserServiceClient interface {
	user.UserServiceClient
}
