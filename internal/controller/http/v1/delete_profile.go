package v1

import (
	"context"

	http_server "github.com/amagkn/golang-production-ready-reference/gen/http/profile_v1/server"
	"github.com/amagkn/golang-production-ready-reference/internal/dto"
	"github.com/amagkn/golang-production-ready-reference/pkg/render"
)

func (h *Handlers) DeleteProfileByID(ctx context.Context, request http_server.DeleteProfileByIDRequestObject,
) (http_server.DeleteProfileByIDResponseObject, error) {
	input := dto.DeleteProfileInput{
		ID: request.ID.String(),
	}

	err := h.usecase.DeleteProfile(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return http_server.DeleteProfileByID400JSONResponse{Error: err.Error()}, nil //nolint:nilerr
	}

	return http_server.DeleteProfileByID204Response{}, nil
}
