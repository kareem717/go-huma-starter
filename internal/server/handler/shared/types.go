package shared

import (
	"github.com/google/uuid"
)

type PaginationRequest struct {
	Cursor int `query:"cursor" required:"false" default:"0"`
	Limit  int `query:"limit" required:"false" default:"10"`
}

type PathIDParam struct {
	ID int `path:"id"`
}

type PathUserIDParam struct {
	UserID uuid.UUID `path:"userId"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type PaginationResponse struct {
	NextCursor *int `json:"nextC˝ursor"`
	HasMore    bool `json:"hasMore"`
}
