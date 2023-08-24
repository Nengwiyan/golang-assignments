package entity

type Pagination struct {
	LastPage int   `json:"last_page"`
	Limit    int   `json:"limit"`
	Offset   int   `json:"offset"`
	Page     int   `json:"page"`
	Total    int64 `json:"total"`
}
