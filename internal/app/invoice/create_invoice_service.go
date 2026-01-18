package invoice

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
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

	if newInvoice.InternalID == "" {
		return nil, go_kit.ErrBadRequest.WithMessage("Internal ID cannot be empty")
	}

	if err := s.repo.Create(ctx, newInvoice); err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, cmd.UserID, newInvoice.ID)
}
