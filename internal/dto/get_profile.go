package dto

import (
	"github.com/amagkn/golang-production-ready-reference/internal/domain"
)

type GetProfileOutput struct {
	domain.Profile
}

type GetProfileInput struct {
	ID string
}
