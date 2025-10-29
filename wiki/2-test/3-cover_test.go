package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Покрытие тестами
// go test -cover ./...

// Покрытие кода тестами
// go test -coverprofile=cover.out ./... && go tool cover -html=cover.out
// В WSL: file://wsl.localhost/Ubuntu/tmp/cover3459369552/coverage.html

func TestValidatePhone(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		wantErr error
	}{
		{
			name:    "Correct phone",
			phone:   "71234567890",
			wantErr: nil,
		},
		{
			name:    "Incorrect length",
			phone:   "7123456789",
			wantErr: ErrLengthPhone,
		},
		{
			name:    "Incorrect start",
			phone:   "81234567890",
			wantErr: ErrPhoneStart,
		},
		{
			name:    "Incorrect digit",
			phone:   "712345PHONE",
			wantErr: ErrPhoneDigit,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.phone)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
