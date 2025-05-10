package requests

type DataWithPaginationRequest struct {
	Query string `json:"query" validate:"required"`
	Limit int    `json:"limit" validate:"required"`
	Page  int    `json:"page" validate:"required"`
}
