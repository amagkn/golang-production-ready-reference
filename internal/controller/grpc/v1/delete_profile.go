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

func (h Handlers) DeleteProfile(ctx context.Context, i *pb.DeleteProfileInput) (*emptypb.Empty, error) {
	input := dto.DeleteProfileInput{
		ID: i.GetId(),
	}

	err := h.usecase.DeleteProfile(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
