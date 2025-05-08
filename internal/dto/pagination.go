package dto

type DataWithPaginationDTI struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

type DataWithPaginationDTO struct {
	Total       int           `json:"total"`
	Limit       int           `json:"limit"`
	CurrentPage int           `json:"currentPage"`
	LastPage    int           `json:"lastPage"`
	Data        []interface{} `json:"data"`
}
