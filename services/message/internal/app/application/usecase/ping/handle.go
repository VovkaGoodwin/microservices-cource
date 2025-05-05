package ping

import "context"

func (u *usecase) Handle(_ context.Context) Response {
	return Response{
		Message: "pong",
	}
}
