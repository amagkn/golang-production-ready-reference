package v1

import "github.com/amagkn/golang-production-ready-reference/internal/usecase"

type Handlers struct {
	usecase *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{
		usecase: uc,
	}
}
