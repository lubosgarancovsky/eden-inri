package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	ClientID      uuid.UUID
	Name          string
	Note          string
	ExternalID    string
	ExternalLink  string
	Total         float64
	BillableHours float64
	IssuedAt      time.Time
	DueAt         time.Time
	DeliveredAt   time.Time
	PaidAt        *time.Time
	IsCanceled    bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewInvoice() (*Invoice, error) {
	// TODO: Implement method
	return nil, errors.New("not implemented")
}
