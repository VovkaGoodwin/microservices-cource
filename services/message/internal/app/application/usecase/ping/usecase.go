package ping

import "context"

type Usecase interface {
	Handle(ctx context.Context) Response
}

var _ Usecase = (*usecase)(nil)

type usecase struct {
}

func New() Usecase {
	return &usecase{}
}
