package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/amagkn/golang-production-ready-reference/gen/grpc/profile_v1"
	"github.com/amagkn/golang-production-ready-reference/internal/dto"
	"github.com/amagkn/golang-production-ready-reference/pkg/render"
)

func (h Handlers) CreateProfile(ctx context.Context, i *pb.CreateProfileInput) (*pb.CreateProfileOutput, error) {
	input := dto.CreateProfileInput{
		Name:  i.GetName(),
		Age:   int(i.GetAge()),
		Email: i.GetEmail(),
		Phone: i.GetPhone(),
	}

	output, err := h.usecase.CreateProfile(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.CreateProfileOutput{
		Id: output.ID.String(),
	}, nil
}
