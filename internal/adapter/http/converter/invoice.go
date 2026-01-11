package converter

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ToCreateInvoiceCommand(input *dto.CreateInvoiceReq) *command.CreateInvoiceCommand {
	return &command.CreateInvoiceCommand{
		UserID:        input.UserID,
		ClientID:      input.ClientID,
		Name:          input.Name,
		Description:   input.Description,
		ExternalID:    input.ExternalID,
		ExternalLink:  input.ExternalLink,
		Total:         input.Total,
		BillableHours: input.BillableHours,
		IssuedAt:      input.IssuedAt,
		DueAt:         input.DueAt,
		DeliveredAt:   input.DeliveredAt,
		PaidAt:        input.PaidAt,
		IsCanceled:    input.IsCanceled,
	}
}

func ToUpdateInvoiceCommand(input *dto.UpdateInvoiceReq) *command.UpdateInvoiceCommand {
	return &command.UpdateInvoiceCommand{
		ID: input.InvoiceID,
		CreateInvoiceCommand: command.CreateInvoiceCommand{
			UserID:        input.UserID,
			ClientID:      input.ClientID,
			Name:          input.Name,
			Description:   input.Description,
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
	}
}

func ToInvoiceResponse(invoice *entity.Invoice) *dto.InvoiceRes {
	return &dto.InvoiceRes{
		ID:            invoice.ID,
		Name:          invoice.Name,
		Description:   invoice.Description,
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
