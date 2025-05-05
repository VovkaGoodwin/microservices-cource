package healthcheck

import "context"

type Interactor interface {
	LivenessProbe(ctx context.Context) LivenessResponse
}

var _ Interactor = (*interactor)(nil)

type interactor struct {
}

func New() Interactor {
	return &interactor{}
}
