package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateInvoiceReq struct {
	UserID        string     `header:"X-User-ID"`
	Name          string     `json:"name"`
	InternalID    string     `json:"internalId"`
	Description   *string    `json:"description"`
	ExternalID    *string    `json:"externalId"`
	ClientID      uuid.UUID  `json:"clientId"`
	Total         float64    `json:"total"`
	BillableHours float64    `json:"billableHours"`
	IssuedAt      time.Time  `json:"issuedAt"`
	DueAt         time.Time  `json:"dueAt"`
	DeliveredAt   time.Time  `json:"deliveredAt"`
	PaidAt        *time.Time `json:"paidAt"`
	IsCanceled    bool       `json:"isCanceled"`
	ExternalLink  *string    `json:"externalLink"`
}

type UpdateInvoiceReq struct {
	UserID    string `header:"X-User-ID"`
	InvoiceID string `uri:"invoiceId"`
	CreateInvoiceReq
}

type DeleteInvoiceReq struct {
	UserID string `header:"X-User-ID"`
	ID     string `uri:"invoiceId"`
}

type FindInvoiceByIDReq struct {
	UserID string `header:"X-User-ID"`
	ID     string `uri:"invoiceId"`
}

type ListInvoicesReq struct {
	UserID string `header:"X-User-ID"`
}

type InvoiceRes struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	InternalID    string     `json:"internalId"`
	Description   *string    `json:"note"`
	ExternalID    *string    `json:"externalId"`
	Total         float64    `json:"total"`
	BillableHours float64    `json:"billableHours"`
	IssuedAt      time.Time  `json:"issuedAt"`
	DueAt         time.Time  `json:"dueAt"`
	DeliveredAt   time.Time  `json:"deliveredAt"`
	PaidAt        *time.Time `json:"paidAt"`
	IsCanceled    bool       `json:"isCanceled"`
	ExternalLink  *string    `json:"externalLink"`
	Client        ClientRes  `json:"client"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
}
