package entity

import "time"

type InvoiceStats struct {
	Count         int64
	Total         float64
	BillableHours float64
}

type InvoiceMonthlyRevenue struct {
	Month   time.Time
	Revenue float64
}
