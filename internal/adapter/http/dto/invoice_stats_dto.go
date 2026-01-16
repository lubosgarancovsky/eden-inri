package dto

type InvoiceStatsReq struct {
	UserID string `header:"X-User-ID"`
}

type InvoiceStatsRes struct {
	Count         int64   `json:"count"`
	Total         float64 `json:"total"`
	BillableHours float64 `json:"billableHours"`
}
