package invoice

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.UpdateInvoiceUseCase = (*UpdateInvoiceService)(nil)

type UpdateInvoiceService struct {
	repo ports.PersistInvoicePort
}

func NewUpdateInvoiceService(repo ports.PersistInvoicePort) *UpdateInvoiceService {
	return &UpdateInvoiceService{repo: repo}
}

func (s *UpdateInvoiceService) Execute(ctx context.Context, cmd *command.UpdateInvoiceCommand) (*entity.Invoice, error) {
	invoice, err := s.repo.FindByID(ctx, cmd.UserID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(invoice)

	if err = s.repo.Update(ctx, invoice); err != nil {
		return nil, err
	}
	return invoice, nil
}
