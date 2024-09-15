package foo

import (
	"proj/internal/entities/shared"

	"github.com/uptrace/bun"
)

// Foo represents an foo entity.
type Foo struct {
	bun.BaseModel `bun:"table:foos"`

	ID   int    `json:"id"`
	Name string `json:"name"`
	shared.Timestamps
}

// CreateFooParams contains the parameters for creating a new foo.
type CreateFooParams struct {
	Name string `json:"name"`
}

// UpdateFooParams contains the parameters for updating a foo.
type UpdateFooParams struct {
	Name string `json:"name"`
}
