package v1

import (
	"context"

	http_server "github.com/amagkn/golang-production-ready-reference/gen/http/profile_v1/server"
	"github.com/amagkn/golang-production-ready-reference/internal/dto"
	"github.com/amagkn/golang-production-ready-reference/pkg/render"
)

func (h *Handlers) CreateProfile(ctx context.Context, request http_server.CreateProfileRequestObject,
) (http_server.CreateProfileResponseObject, error) {
	input := dto.CreateProfileInput{
		Name:  request.Body.Name,
		Age:   request.Body.Age,
		Email: request.Body.Email,
		Phone: request.Body.Phone,
	}

	output, err := h.usecase.CreateProfile(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return http_server.CreateProfile400JSONResponse{Error: err.Error()}, nil
	}

	return http_server.CreateProfile200JSONResponse{
		ID: output.ID,
	}, nil
}
