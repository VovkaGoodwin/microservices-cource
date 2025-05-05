package interactors

import (
	"context"

	"gateway/internal/app/application/dto"
	"gateway/proto/message"
)

type Message struct {
	client message.MessageServiceClient
}

func NewMessage(client message.MessageServiceClient) *Message {
	return &Message{
		client: client,
	}
}

func (m *Message) Ping(ctx context.Context) (string, error) {
	result, err := m.client.Ping(ctx, &message.PingRequest{})
	if err != nil {
		return "", err
	}

	return result.Message, nil
}

func (m *Message) GetMessagesByUserId(ctx context.Context, userId string) (dto.GetMessagesByUSerIdResponse, error) {
	result, err := m.client.GetMessagesByUserId(ctx, &message.GetMessagesByUserIdRequest{UserId: userId})
	if err != nil {
		return dto.GetMessagesByUSerIdResponse{}, err
	}

	var messages []dto.Message
	for _, msg := range result.Messages {
		item := dto.Message{
			Id:   msg.Id,
			Text: msg.Text,
		}

		messages = append(messages, item)
	}

	return dto.GetMessagesByUSerIdResponse{Messages: messages}, nil
}

func (m *Message) SendMessage(ctx context.Context, sendMessageDto dto.SendMessageRequest) (dto.SendMessageResponse, error) {
	result, err := m.client.SendMessage(ctx, &message.SendMessageRequest{
		UserId: sendMessageDto.UserId,
		Text:   sendMessageDto.Text,
	})
	if err != nil {
		return dto.SendMessageResponse{}, err
	}

	return dto.SendMessageResponse{Success: result.Success}, nil
}
