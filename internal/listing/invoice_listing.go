package listing

var InvoiceFilter = map[string]string{
	"externalId":    "external_id",
	"clientId":      "client_id",
	"isCanceled":    "is_canceled",
	"billableHours": "billable_hours",
	"total":         "total",
	"issuedAt":      "issued_at",
	"dueAt":         "due_at",
	"payedAt":       "payed_at",
}

var InvoiceSort = map[string]string{
	"issuedAt":      "issued_at",
	"dueAt":         "due_at",
	"createdAt":     "created_at",
	"billableHours": "billable_hours",
	"total":         "total",
	"payedAt":       "payed_at",
}
