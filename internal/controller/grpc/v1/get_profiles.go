package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/amagkn/golang-production-ready-reference/gen/grpc/profile_v1"
	"github.com/amagkn/golang-production-ready-reference/internal/dto"
	"github.com/amagkn/golang-production-ready-reference/pkg/render"
)

func (h Handlers) GetProfiles(ctx context.Context, i *pb.GetProfilesInput) (*pb.GetProfilesOutput, error) {
	input := dto.GetProfilesInput{
		Sort: i.GetSort(),
	}

	if i.Order != nil {
		input.Order = i.GetOrder()
	}

	if i.Limit != nil {
		input.Limit = int(i.GetLimit())
	}

	if i.Offset != nil {
		input.Offset = int(i.GetOffset())
	}

	o, err := h.usecase.GetProfiles(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	profiles := make([]*pb.GetProfilesOutput_Profile, 0, len(o.Profiles))

	for _, profile := range o.Profiles {
		p := &pb.GetProfilesOutput_Profile{
			Id:        profile.ID.String(),
			CreatedAt: timestamppb.New(profile.CreatedAt),
			UpdatedAt: timestamppb.New(profile.UpdatedAt),
			Name:      string(profile.Name),
			Age:       int32(profile.Age),
			Status:    int32(profile.Status),
			Verified:  profile.Verified,
			Contacts: &pb.GetProfilesOutput_Profile_Contacts{
				Email: profile.Contacts.Email,
				Phone: profile.Contacts.Phone,
			},
		}

		profiles = append(profiles, p)
	}

	return &pb.GetProfilesOutput{
		Profiles: profiles,
	}, nil
}
