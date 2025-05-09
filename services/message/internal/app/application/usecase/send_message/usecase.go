package send_message

import "context"

type Usecase interface {
	Handle(ctx context.Context, request Request) (Response, error)
}

var _ Usecase = (*usecase)(nil)

type usecase struct {
}

func New() Usecase {
	return &usecase{}
}
