package usecase

import "auth/internal/app/application/dto"

type Healthcheck struct{}

func NewHealthcheckUseCase() *Healthcheck {
	return &Healthcheck{}
}

func (h *Healthcheck) Handle() dto.HealthcheckResponseDto {
	return dto.HealthcheckResponseDto{
		Result: "OK",
	}
}
