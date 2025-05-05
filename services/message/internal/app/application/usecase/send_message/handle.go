package send_message

import "context"

func (u *usecase) Handle(ctx context.Context, req Request) (Response, error) {
	return Response{
		Success: true,
	}, nil
}
