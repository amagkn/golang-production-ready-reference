package profile_client_grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	pb "github.com/amagkn/golang-production-ready-reference/gen/grpc/profile_v1"
)

func (c *Client) Update(ctx context.Context, id string, name *string, age *int, email, phone *string) error {
	input := &pb.UpdateProfileInput{
		Id:             id,
		Name:           name,
		Age:            parseAge(age),
		Email:          email,
		Phone:          phone,
		IdempotencyKey: uuid.New().String(),
	}

	_, err := c.client.UpdateProfile(ctx, input)
	if err != nil {
		return fmt.Errorf("c.client.UpdateProfile: %w", err)
	}

	return nil
}

func parseAge(age *int) *int32 {
	if age == nil {
		return nil
	}

	a := int32(*age) //nolint:gosec

	return &a
}
