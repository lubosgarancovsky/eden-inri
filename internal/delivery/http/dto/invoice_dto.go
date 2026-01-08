package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateInvoiceReq struct {
	Name          string     `json:"name" binding:"required,min=3,max=100"`
	Note          string     `json:"note"`
	ExternalID    string     `json:"externalId"`
	ClientID      uuid.UUID  `json:"clientId" binding:"required"`
	Total         float64    `json:"total" binding:"required, min=0"`
	BillableHours float64    `json:"billableHours" binding:"required,min=0"`
	IssuedAt      time.Time  `json:"issuedAt" binding:"required"`
	DueAt         time.Time  `json:"dueAt" binding:"required"`
	DeliveredAt   time.Time  `json:"deliveredAt" binding:"required"`
	PaidAt        *time.Time `json:"paidAt"`
	IsCanceled    bool       `json:"isCanceled"`
	ExternalLink  string     `json:"externalLink"`
}

type UpdateInvoiceReq struct {
	InvoiceID uuid.UUID `uri:"invoiceId" binding:"required"`
	CreateInvoiceReq
}

type InvoiceRes struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Note          string     `json:"note"`
	ExternalID    string     `json:"externalId"`
	Total         float64    `json:"total"`
	BillableHours float64    `json:"billableHours"`
	IssuedAt      time.Time  `json:"issuedAt"`
	DueAt         time.Time  `json:"dueAt"`
	DeliveredAt   time.Time  `json:"deliveredAt"`
	PaidAt        *time.Time `json:"paidAt"`
	IsCanceled    bool       `json:"isCanceled"`
	ExternalLink  string     `json:"externalLink"`
	Client        ClientRes  `json:"client"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
}
