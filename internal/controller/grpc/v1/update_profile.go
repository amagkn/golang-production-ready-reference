package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/amagkn/golang-production-ready-reference/gen/grpc/profile_v1"
	"github.com/amagkn/golang-production-ready-reference/internal/dto"
	"github.com/amagkn/golang-production-ready-reference/pkg/render"
)

func (h Handlers) UpdateProfile(ctx context.Context, i *pb.UpdateProfileInput) (*emptypb.Empty, error) {
	input := dto.UpdateProfileInput{
		ID:    i.GetId(),
		Name:  i.Name,
		Age:   parseAge(i.Age),
		Email: i.Email,
		Phone: i.Phone,
	}

	err := h.usecase.UpdateProfile(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func parseAge(age *int32) *int {
	if age == nil {
		return nil
	}

	a := int(*age)

	return &a
}
