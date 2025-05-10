package dto

type DataWithPaginationDTO struct {
	Total       int           `json:"total"`
	Limit       int           `json:"limit"`
	CurrentPage int           `json:"currentPage"`
	LastPage    int           `json:"lastPage"`
	Data        []interface{} `json:"data"`
}
