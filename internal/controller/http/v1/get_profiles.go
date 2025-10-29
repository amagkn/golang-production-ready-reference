package v1

import (
	"context"

	http_server "github.com/amagkn/golang-production-ready-reference/gen/http/profile_v1/server"
	"github.com/amagkn/golang-production-ready-reference/internal/dto"
	"github.com/amagkn/golang-production-ready-reference/pkg/render"
)

func (h *Handlers) GetProfiles(ctx context.Context, request http_server.GetProfilesRequestObject,
) (http_server.GetProfilesResponseObject, error) {
	input := dto.GetProfilesInput{
		Sort: request.Params.Sort,
	}

	if request.Params.Order != nil {
		input.Order = *request.Params.Order
	}

	if request.Params.Offset != nil {
		input.Offset = *request.Params.Offset
	}

	if request.Params.Limit != nil {
		input.Limit = *request.Params.Limit
	}

	output, err := h.usecase.GetProfiles(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return http_server.GetProfiles400JSONResponse{Error: err.Error()}, nil
	}

	profiles := make(http_server.GetProfiles200JSONResponse, 0, len(output.Profiles))

	for _, profile := range output.Profiles {
		var p http_server.GetProfileOutput

		p.ID = profile.ID
		p.CreatedAt = profile.CreatedAt
		p.UpdatedAt = profile.UpdatedAt
		p.Name = string(profile.Name)
		p.Age = int(profile.Age)
		p.Status = int(profile.Status)
		p.Verified = profile.Verified
		p.Contacts.Email = profile.Contacts.Email
		p.Contacts.Phone = profile.Contacts.Phone

		profiles = append(profiles, p)
	}

	return profiles, nil
}
