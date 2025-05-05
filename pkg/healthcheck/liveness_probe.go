package healthcheck

import "context"

func (i *interactor) LivenessProbe(_ context.Context) LivenessResponseDto {
	return LivenessResponseDto{
		Result: "OK",
	}
}
