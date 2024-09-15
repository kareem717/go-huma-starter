package shared

import (
	"time"
)

// Timestamps contains the timestamps for an entity.
type Timestamps struct {
	// CreatedAt is the timestamp of the creation.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt is the timestamp of the last update. Null until the first update.
	UpdatedAt *time.Time `json:"updatedAt"`
	// DeletedAt is the timestamp of the deletion. Null until the entity is deleted.
	DeletedAt *time.Time `json:"deletedAt" bun:",soft_delete"`
}
