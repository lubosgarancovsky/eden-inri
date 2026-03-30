package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateTaxCommand struct {
	UserID      uuid.UUID
	Category    entity.TaxCategory
	Amount      float64
	Currency    string
	Description *string
	Reference   *string
	Period      *time.Time
	PaidAt      *time.Time
}

func (c *CreateTaxCommand) ToDomain() *entity.Tax {
	return &entity.Tax{
		ID:          uuid.New(),
		UserID:      c.UserID,
		Category:    c.Category,
		Amount:      c.Amount,
		Currency:    c.Currency,
		Description: c.Description,
		Reference:   c.Reference,
		Period:      c.Period,
		PaidAt:      c.PaidAt,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type UpdateTaxCommand struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Description *string
	Reference   *string
	PaidAt      *time.Time
}

func (c *UpdateTaxCommand) Apply(e *entity.Tax) {
	e.Description = c.Description
	e.Reference = c.Reference
	e.PaidAt = c.PaidAt
}

type DeleteTaxCommand struct {
	ID     uuid.UUID
	UserID uuid.UUID
}
