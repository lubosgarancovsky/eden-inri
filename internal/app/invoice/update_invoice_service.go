package invoice

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

var _ ports.UpdateInvoiceUseCase = (*UpdateInvoiceService)(nil)

type UpdateInvoiceService struct {
	repo ports.PersistInvoicePort
	tm   ports.TransactionManager
}

func NewUpdateInvoiceService(repo ports.PersistInvoicePort, tm ports.TransactionManager) *UpdateInvoiceService {
	return &UpdateInvoiceService{repo: repo, tm: tm}
}

func (s *UpdateInvoiceService) Execute(ctx context.Context, inv *ports.Invoice) (*ports.Invoice, error) {
	if err := s.tm.WithTransaction(ctx, func(ctx context.Context) error { return s.repo.Update(ctx, inv) }); err != nil {
		return nil, err
	}
	return inv, nil
}
