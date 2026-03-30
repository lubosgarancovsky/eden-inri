package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
)

func ToCreateInvoiceCommand(input *dto.CreateInvoiceReq) (*command.CreateInvoiceCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateInvoiceCommand{
		UserID:        userID,
		ClientID:      input.ClientID,
		Name:          input.Name,
		Description:   input.Description,
		ExternalID:    input.ExternalID,
		InternalID:    input.InternalID,
		ExternalLink:  input.ExternalLink,
		Total:         input.Total,
		BillableHours: input.BillableHours,
		IssuedAt:      input.IssuedAt,
		DueAt:         input.DueAt,
		DeliveredAt:   input.DeliveredAt,
		PaidAt:        input.PaidAt,
		IsCanceled:    input.IsCanceled,
	}, nil
}

func ToUpdateInvoiceCommand(input *dto.UpdateInvoiceReq) (*command.UpdateInvoiceCommand, error) {
	id, err := uuid.Parse(input.InvoiceID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateInvoiceCommand{
		ID: id,
		CreateInvoiceCommand: command.CreateInvoiceCommand{
			UserID:        userID,
			ClientID:      input.ClientID,
			Name:          input.Name,
			Description:   input.Description,
			InternalID:    input.InternalID,
			ExternalID:    input.ExternalID,
			ExternalLink:  input.ExternalLink,
			Total:         input.Total,
			BillableHours: input.BillableHours,
			IssuedAt:      input.IssuedAt,
			DueAt:         input.DueAt,
			DeliveredAt:   input.DeliveredAt,
			PaidAt:        input.PaidAt,
			IsCanceled:    input.IsCanceled,
		},
	}, nil
}

func ToInvoiceResponse(invoice *entity.Invoice) *dto.InvoiceRes {
	return &dto.InvoiceRes{
		ID:            invoice.ID,
		Name:          invoice.Name,
		Description:   invoice.Description,
		InternalID:    invoice.InternalID,
		ExternalID:    invoice.ExternalID,
		Total:         invoice.Total,
		BillableHours: invoice.BillableHours,
		IssuedAt:      invoice.IssuedAt,
		DueAt:         invoice.DueAt,
		DeliveredAt:   invoice.DeliveredAt,
		PaidAt:        invoice.PaidAt,
		IsCanceled:    invoice.IsCanceled,
		ExternalLink:  invoice.ExternalLink,
		CreatedAt:     invoice.CreatedAt,
		UpdatedAt:     invoice.UpdatedAt,
		Client:        *ToClientResponse(invoice.Client),
	}
}
