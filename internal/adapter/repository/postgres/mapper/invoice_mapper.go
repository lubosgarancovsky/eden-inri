package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func InvoiceFromDomain(entity *entity.Invoice) *model.Invoice {
	return &model.Invoice{
		ID:            entity.ID,
		UserID:        entity.UserID,
		ClientID:      entity.ClientID,
		Name:          entity.Name,
		Description:   entity.Description,
		ExternalID:    entity.ExternalID,
		ExternalLink:  entity.ExternalLink,
		Total:         entity.Total,
		BillableHours: entity.BillableHours,
		IsCanceled:    entity.IsCanceled,
		IssuedAt:      entity.IssuedAt,
		DueAt:         entity.DueAt,
		DeliveredAt:   entity.DeliveredAt,
		PaidAt:        entity.PaidAt,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
		Client:        *ClientFromDomain(entity.Client),
	}
}
