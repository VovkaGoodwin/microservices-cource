package healthcheck

type ResponseDto struct {
	Result string
}

type LivenessResponseDto struct {
	Result string
}

type ReadinessResponseDto struct {
	Result  string
	Success bool
}
