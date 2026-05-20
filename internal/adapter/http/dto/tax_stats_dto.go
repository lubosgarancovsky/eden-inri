package dto

type TaxStatsReq struct {
	UserID string `header:"X-User-ID"`
}

type TaxStatsRes struct {
	Count int64   `json:"count"`
	Total float64 `json:"total"`
}
