package invoice

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.CreateInvoiceUseCase = (*CreateInvoiceService)(nil)

type CreateInvoiceService struct {
	repo ports.PersistInvoicePort
}

func NewCreateInvoiceService(repo ports.PersistInvoicePort) *CreateInvoiceService {
	return &CreateInvoiceService{repo: repo}
}

func (s *CreateInvoiceService) Execute(ctx context.Context, cmd *command.CreateInvoiceCommand) (*entity.Invoice, error) {
	newInvoice := cmd.ToDomain()
	if err := s.repo.Create(ctx, newInvoice); err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, cmd.UserID, newInvoice.ID)
}
