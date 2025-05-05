package clients

import "github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/message"

type MessageServiceClient interface {
	message.MessageServiceClient
}
