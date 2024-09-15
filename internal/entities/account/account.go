package account

import (
	"proj/internal/entities/shared"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Account represents an account entity.
type Account struct {
	bun.BaseModel `bun:"table:accounts"`

	ID     int       `json:"id"`
	UserID uuid.UUID `json:"userId" format:"uuid" minLength:"36" maxLength:"36"`
	Name   string    `json:"name"`
	shared.Timestamps
}

// CreateAccountParams contains the parameters for creating a new account.
type CreateAccountParams struct {
	Name   string    `json:"name"`
	UserID uuid.UUID `json:"userId" format:"uuid" minLength:"36" maxLength:"36"`
}

// UpdateAccountParams contains the parameters for updating a account.
type UpdateAccountParams struct {
	Name string `json:"name"`
}
