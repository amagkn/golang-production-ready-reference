package profile_client_gen

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	http_client "github.com/amagkn/golang-production-ready-reference/gen/http/profile_v1/client"
)

func (c *Client) Update(ctx context.Context, id string, name *string, age *int, email, phone *string) error {
	input := http_client.UpdateProfileInput{
		ID:             uuid.MustParse(id),
		Name:           name,
		Age:            age,
		Email:          email,
		Phone:          phone,
		IdempotencyKey: uuid.New().String(),
	}

	output, err := c.client.UpdateProfileWithResponse(ctx, input)
	if err != nil {
		return fmt.Errorf("delete profile: %w", err)
	}

	if output.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("request failed: status: %s, body:%s", output.Status(), output.Body)
	}

	return nil
}
