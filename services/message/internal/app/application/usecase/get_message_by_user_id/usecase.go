package get_message_by_user_id

import "context"

type Usecase interface {
	Handle(ctx context.Context, request GetMessageByUserIdRequest) (GetMessageByUserIdResponse, error)
}

var _ Usecase = (*usecase)(nil)

type usecase struct {
}

func New() Usecase {
	return &usecase{}
}
