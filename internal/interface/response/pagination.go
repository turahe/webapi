package response

type PaginationResponse struct {
	Data         interface{} `json:"data"`
	TotalCount   int         `json:"total_count"`
	TotalPage    int         `json:"total_page"`
	CurrentPage  int         `json:"current_page"`
	LastPage     int         `json:"last_page"`
	PerPage      int         `json:"per_page"`
	NextPage     int         `json:"next_page"`
	PreviousPage int         `json:"previous_page"`
	Path         string      `json:"path"`
}
