package invoice

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

var _ ports.CreateInvoiceUseCase = (*CreateInvoiceService)(nil)

type CreateInvoiceService struct {
	repo ports.PersistInvoicePort
	tm   ports.TransactionManager
}

func NewCreateInvoiceService(repo ports.PersistInvoicePort, tm ports.TransactionManager) *CreateInvoiceService {
	return &CreateInvoiceService{repo: repo, tm: tm}
}

func (s *CreateInvoiceService) Execute(ctx context.Context, inv *ports.Invoice) (*ports.Invoice, error) {
	if err := s.tm.WithTransaction(ctx, func(ctx context.Context) error { return s.repo.Create(ctx, inv) }); err != nil {
		return nil, err
	}
	return inv, nil
}
