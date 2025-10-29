package profile_client_grpc

import (
	"context"
	"fmt"

	pb "github.com/amagkn/golang-production-ready-reference/gen/grpc/profile_v1"
	"github.com/amagkn/golang-production-ready-reference/pkg/profile_client"
)

type Profile profile_client.Profile

func (c *Client) GetProfile(ctx context.Context, id string) (Profile, error) {
	input := &pb.GetProfileInput{
		Id: id,
	}

	o, err := c.client.GetProfile(ctx, input)
	if err != nil {
		return Profile{}, fmt.Errorf("client.Get: %w", err)
	}

	return Profile{
		ID:        o.GetId(),
		CreatedAt: o.GetCreatedAt().String(),
		UpdatedAt: o.GetUpdatedAt().String(),
		Name:      o.GetName(),
		Age:       int(o.GetAge()),
		Status:    int(o.GetStatus()),
		Verified:  o.GetVerified(),
		Contacts: struct {
			Email string `json:"email"`
			Phone string `json:"phone"`
		}{
			Email: o.GetContacts().GetEmail(),
			Phone: o.GetContacts().GetPhone(),
		},
	}, nil
}
