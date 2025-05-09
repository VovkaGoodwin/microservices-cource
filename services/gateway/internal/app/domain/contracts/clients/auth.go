package clients

import "github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/auth"

type AuthServiceClient interface {
	auth.AuthServiceClient
}
