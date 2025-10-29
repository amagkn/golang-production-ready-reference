package v1

import (
	pb "github.com/amagkn/golang-production-ready-reference/gen/grpc/profile_v1"
	"github.com/amagkn/golang-production-ready-reference/internal/usecase"
)

type Handlers struct {
	pb.UnimplementedProfileV1Server
	usecase *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{
		usecase: uc,
	}
}
