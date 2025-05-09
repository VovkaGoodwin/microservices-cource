package message

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/message"
)

func (m *interactor) Ping(ctx context.Context) (string, error) {
	result, err := m.client.Ping(ctx, &message.PingRequest{})
	if err != nil {
		return "", err
	}

	return result.Message, nil
}
