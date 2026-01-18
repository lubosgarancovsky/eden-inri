package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateInvoiceCommand struct {
	UserID        uuid.UUID
	ClientID      uuid.UUID
	Name          string
	InternalID    string
	Description   *string
	ExternalID    *string
	ExternalLink  *string
	Total         float64
	BillableHours float64
	IssuedAt      time.Time
	DueAt         time.Time
	DeliveredAt   time.Time
	PaidAt        *time.Time
	IsCanceled    bool
}

func (c *CreateInvoiceCommand) ToDomain() *entity.Invoice {
	now := time.Now()
	return &entity.Invoice{
		ID:            uuid.New(),
		UserID:        c.UserID,
		ClientID:      c.ClientID,
		Name:          c.Name,
		InternalID:    c.InternalID,
		Description:   c.Description,
		ExternalID:    c.ExternalID,
		ExternalLink:  c.ExternalLink,
		Total:         c.Total,
		BillableHours: c.BillableHours,
		IssuedAt:      c.IssuedAt,
		DueAt:         c.DueAt,
		DeliveredAt:   c.DeliveredAt,
		PaidAt:        c.PaidAt,
		IsCanceled:    c.IsCanceled,
		CreatedAt:     now,
		UpdatedAt:     now,
		Client:        nil,
	}
}

type UpdateInvoiceCommand struct {
	ID uuid.UUID
	CreateInvoiceCommand
}

func (c *UpdateInvoiceCommand) Apply(inv *entity.Invoice) {
	inv.Name = c.Name
	inv.ClientID = c.ClientID
	inv.Description = c.Description
	inv.InternalID = c.InternalID
	inv.ExternalID = c.ExternalID
	inv.ExternalLink = c.ExternalLink
	inv.Total = c.Total
	inv.BillableHours = c.BillableHours
	inv.IssuedAt = c.IssuedAt
	inv.DueAt = c.DueAt
	inv.DeliveredAt = c.DeliveredAt
	inv.PaidAt = c.PaidAt
	inv.IsCanceled = c.IsCanceled
	inv.UpdatedAt = time.Now()
}

type DeleteInvoiceCommand struct {
	ID     uuid.UUID
	UserID uuid.UUID
}
