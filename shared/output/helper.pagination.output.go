package opt

type Pagination struct {
	Page      int     `json:"page"`
	Limit     int     `json:"per_page"`
	TotalPage float64 `json:"total_page"`
	TotalData int     `json:"total_data"`
}
