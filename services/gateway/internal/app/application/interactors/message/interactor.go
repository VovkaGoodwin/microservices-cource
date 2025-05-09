package message

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/domain/contracts/clients"
)

type Interactor interface {
	GetMessagesByUserId(ctx context.Context, userId string) (GetMessagesByUSerIdResponse, error)
	Ping(ctx context.Context) (string, error)
	SendMessage(ctx context.Context, sendMessageDto SendMessageRequest) (SendMessageResponse, error)
}

type Deps struct {
	client clients.MessageServiceClient
}

type interactor struct {
	*Deps
}

func New(deps *Deps) Interactor {
	return &interactor{
		deps,
	}
}
