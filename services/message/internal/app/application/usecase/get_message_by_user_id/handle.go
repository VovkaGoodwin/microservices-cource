package get_message_by_user_id

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

func (u *usecase) Handle(_ context.Context, _ GetMessageByUserIdRequest) (GetMessageByUserIdResponse, error) {
	var messages []Message
	for i := 0; i < 10; i++ {
		item := Message{
			Id:   uuid.New().String(),
			Text: strings.Repeat("text", i),
		}
		messages = append(messages, item)
	}

	return GetMessageByUserIdResponse{Messages: messages}, nil
}
