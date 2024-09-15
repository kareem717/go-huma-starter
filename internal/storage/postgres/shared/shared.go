package shared

type PaginationRequest struct {
	Cursor int `json:"cursor" default:"1" min:"1" required:"false"`
	Limit  int `json:"limit" default:"10" min:"1" max:"100" required:"false"`
}
