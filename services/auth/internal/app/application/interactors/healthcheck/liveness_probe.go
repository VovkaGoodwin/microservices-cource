package healthcheck

import "context"

func (h *Interactor) LivenessProbe(_ context.Context) LivenessResponseDto {
	return LivenessResponseDto{
		Result: "OK",
	}
}
