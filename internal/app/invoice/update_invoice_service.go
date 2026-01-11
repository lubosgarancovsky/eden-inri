package invoice

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.UpdateInvoiceUseCase = (*UpdateInvoiceService)(nil)

type UpdateInvoiceService struct {
	repository ports.PersistInvoicePort
}

func NewUpdateInvoiceService(repository ports.PersistInvoicePort) *UpdateInvoiceService {
	return &UpdateInvoiceService{repository}
}

func (s *UpdateInvoiceService) Execute(ctx context.Context, cmd *command.UpdateInvoiceCommand) (*entity.Invoice, error) {
	invoice, err := s.repository.FindByID(ctx, cmd.UserID, cmd.ID)
	if err != nil {
		return nil, err
	}

	prevClient := invoice.ClientID

	cmd.Apply(invoice)
	if prevClient != invoice.ClientID {

	}

	if err = s.repository.Update(ctx, invoice); err != nil {
		return nil, err
	}
	return invoice, nil
}
