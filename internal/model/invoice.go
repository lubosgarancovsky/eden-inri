package model

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID            uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name          string     `json:"name"`
	Note          string     `json:"note"`
	ExternalID    string     `json:"externalId"`
	UserID        uuid.UUID  `json:"-"`
	ClientID      uuid.UUID  `json:"-"`
	Client        Client     `json:"client"`
	Total         float64    `json:"total"`
	BillableHours float64    `json:"billableHours"`
	IssuedAt      time.Time  `json:"issuedAt"`
	DueAt         time.Time  `json:"dueAt"`
	DeliveredAt   time.Time  `json:"deliveredAt"`
	PaidAt        *time.Time `json:"paidAt"`
	IsCanceled    bool       `json:"isCanceled"`
	ExternalLink  string     `json:"externalLink"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
}

type InvoiceRequest struct {
	Name          string     `json:"name"`
	Note          string     `json:"note"`
	ExternalID    string     `json:"externalId"`
	ClientID      uuid.UUID  `json:"clientId"`
	Total         float64    `json:"total"`
	BillableHours float64    `json:"billableHours"`
	IssuedAt      time.Time  `json:"issuedAt"`
	DueAt         time.Time  `json:"dueAt"`
	DeliveredAt   time.Time  `json:"deliveredAt"`
	PaidAt        *time.Time `json:"paidAt"`
	IsCanceled    bool       `json:"isCanceled"`
	ExternalLink  string     `json:"externalLink"`
}

func (i *Invoice) TableName() string {
	return "inri_invoice"
}
