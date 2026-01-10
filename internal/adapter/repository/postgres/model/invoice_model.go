package model

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"time"
)

type Invoice struct {
	ID            uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID        uuid.UUID
	ClientID      uuid.UUID
	Number        string
	IssuedAt      time.Time
	DueAt         time.Time
	TotalAmount   int64
	Currency      string
	Status        string
	Note          string
	ExternalID    string
	IsCanceled    bool
	BillableHours int64
	PaidAt        *time.Time
	CreatedAt     time.Time
}

func (Invoice) TableName() string { return "inri_invoices" }

func (m *Invoice) ToPort() *ports.Invoice {
	inv := &ports.Invoice{
		ID: m.ID, UserID: m.UserID, ClientID: m.ClientID, Number: m.Number,
		IssuedAt: m.IssuedAt, DueAt: m.DueAt, TotalAmount: m.TotalAmount,
		Currency: m.Currency, Status: m.Status, Note: m.Note,
	}
	return inv
}

func InvoiceFromPort(p *ports.Invoice) *Invoice {
	return &Invoice{
		ID: p.ID, UserID: p.UserID, ClientID: p.ClientID, Number: p.Number,
		IssuedAt: p.IssuedAt, DueAt: p.DueAt, TotalAmount: p.TotalAmount,
		Currency: p.Currency, Status: p.Status, Note: p.Note,
	}
}
