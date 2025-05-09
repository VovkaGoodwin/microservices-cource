package message

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/message"
)

func (m *interactor) SendMessage(ctx context.Context, sendMessageDto SendMessageRequest) (SendMessageResponse, error) {
	result, err := m.client.SendMessage(ctx, &message.SendMessageRequest{
		UserId: sendMessageDto.UserId,
		Text:   sendMessageDto.Text,
	})
	if err != nil {
		return SendMessageResponse{}, err
	}

	return SendMessageResponse{Success: result.Success}, nil
}
