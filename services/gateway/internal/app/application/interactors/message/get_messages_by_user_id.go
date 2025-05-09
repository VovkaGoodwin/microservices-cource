package message

import (
	"context"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/proto/message"
)

func (m *interactor) GetMessagesByUserId(ctx context.Context, userId string) (GetMessagesByUSerIdResponse, error) {
	result, err := m.client.GetMessagesByUserId(ctx, &message.GetMessagesByUserIdRequest{UserId: userId})
	if err != nil {
		return GetMessagesByUSerIdResponse{}, err
	}

	var messages []Message
	for _, msg := range result.Messages {
		item := Message{
			Id:   msg.Id,
			Text: msg.Text,
		}

		messages = append(messages, item)
	}

	return GetMessagesByUSerIdResponse{Messages: messages}, nil
}
