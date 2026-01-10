package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateClientCommand struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	ClientType   string
	ContractType string
	Name         string
	Description  *string
	TaxNumber    *string
	Address      *string
	Tags         []string
	HourRate     float64
	StartedAt    time.Time
	FinishedAt   *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (c *CreateClientCommand) ToDomain() *entity.Client {
	return &entity.Client{
		ID:           c.ID,
		UserID:       c.UserID,
		ClientType:   entity.ClientType(c.ClientType),
		ContractType: entity.ClientContractType(c.ContractType),
		Name:         c.Name,
		Description:  c.Description,
		TaxNumber:    c.TaxNumber,
		Address:      c.Address,
		Tags:         c.Tags,
		HourRate:     c.HourRate,
		StartedAt:    c.StartedAt,
		FinishedAt:   c.FinishedAt,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}
