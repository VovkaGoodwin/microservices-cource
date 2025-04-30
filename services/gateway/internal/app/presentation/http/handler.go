package http

import "gateway/internal/app/application/interactors"

type Handler struct {
	auth    *interactors.Auth
	user    *interactors.User
	message *interactors.Message
}

func NewHandler(
	auth *interactors.Auth,
	user *interactors.User,
	message *interactors.Message,
) *Handler {
	return &Handler{
		auth:    auth,
		user:    user,
		message: message,
	}
}
