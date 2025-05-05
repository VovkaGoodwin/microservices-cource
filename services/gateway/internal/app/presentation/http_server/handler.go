package http_server

import (
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/application/interactors/auth"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/application/interactors/message"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app/application/interactors/user"
)

type Handler struct {
	auth    auth.Interactor
	user    user.Interactor
	message message.Interactor
}

func NewHandler(
	auth auth.Interactor,
	user user.Interactor,
	message message.Interactor,
) *Handler {
	return &Handler{
		auth:    auth,
		user:    user,
		message: message,
	}
}
