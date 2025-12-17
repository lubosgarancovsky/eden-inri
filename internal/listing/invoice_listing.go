package listing

var InvoiceFilter = map[string]string{
	"externalId":    "external_id",
	"clientId":      "client_id",
	"isCanceled":    "is_canceled",
	"billableHours": "billable_hours",
	"total":         "total",
	"issuedAt":      "issued_at",
	"dueAt":         "due_at",
	"paidAt":        "paid_at",
	"name":          "name",
}

var InvoiceSort = map[string]string{
	"issuedAt":      "issued_at",
	"dueAt":         "due_at",
	"createdAt":     "created_at",
	"billableHours": "billable_hours",
	"total":         "total",
	"paidAt":        "paid_at",
	"name":          "name",
}
