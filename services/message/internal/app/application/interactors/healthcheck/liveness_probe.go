package healthcheck

import "context"

func (i *interactor) LivenessProbe(_ context.Context) LivenessResponse {
	return LivenessResponse{
		Message: "pong",
	}
}
