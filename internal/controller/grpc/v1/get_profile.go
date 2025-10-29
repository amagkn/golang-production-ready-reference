package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/amagkn/golang-production-ready-reference/gen/grpc/profile_v1"
	"github.com/amagkn/golang-production-ready-reference/internal/domain"
	"github.com/amagkn/golang-production-ready-reference/internal/dto"
	"github.com/amagkn/golang-production-ready-reference/pkg/render"
)

func (h Handlers) GetProfile(ctx context.Context, i *pb.GetProfileInput) (*pb.GetProfileOutput, error) {
	input := dto.GetProfileInput{
		ID: i.GetId(),
	}

	o, err := h.usecase.GetProfile(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		switch {
		case errors.Is(err, domain.ErrNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	return &pb.GetProfileOutput{
		Id:        o.ID.String(),
		CreatedAt: timestamppb.New(o.CreatedAt),
		UpdatedAt: timestamppb.New(o.UpdatedAt),
		Name:      string(o.Name),
		Age:       int32(o.Age), //nolint:gosec
		Verified:  o.Verified,
		Status:    int32(o.Status), //nolint:gosec
		Contacts: &pb.GetProfileOutput_Contacts{
			Email: o.Contacts.Email,
			Phone: o.Contacts.Phone,
		},
	}, nil
}
