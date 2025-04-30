package interactors

import (
	"context"
	"gateway/internal/app/application/dto"
	proto "gateway/proto/message"
)

type Message struct {
	client proto.MessageServiceClient
}

func NewMessage(client proto.MessageServiceClient) *Message {
	return &Message{
		client: client,
	}
}

func (m *Message) Ping(ctx context.Context) (string, error) {
	result, err := m.client.Ping(ctx, &proto.PingRequest{})
	if err != nil {
		return "", err
	}

	return result.Message, nil
}

func (m *Message) GetMessagesByUserId(ctx context.Context, userId string) (dto.GetMessagesByUSerIdResponse, error) {
	result, err := m.client.GetMessagesByUserId(ctx, &proto.GetMessagesByUserIdRequest{UserId: userId})
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

func (m *Message) SendMessage(ctx context.Context, message dto.SendMessageRequest) (dto.SendMessageResponse, error) {
	result, err := m.client.SendMessage(ctx, &proto.SendMessageRequest{
		UserId: message.UserId,
		Text:   message.Text,
	})
	if err != nil {
		return dto.SendMessageResponse{}, err
	}

	return dto.SendMessageResponse{Success: result.Success}, nil
}
