package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/amagkn/golang-production-ready-reference/internal/domain"
	"github.com/amagkn/golang-production-ready-reference/internal/dto"
	"github.com/amagkn/golang-production-ready-reference/pkg/otel/tracer"
)

func (u *UseCase) DeleteProfile(ctx context.Context, input dto.DeleteProfileInput) error {
	ctx, span := tracer.Start(ctx, "usecase DeleteProfile")
	defer span.End()

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return domain.ErrUUIDInvalid
	}

	err = u.postgres.DeleteProfile(ctx, id)
	if err != nil {
		return fmt.Errorf("postgres.DeleteProfile: %w", err)
	}

	return nil
}
