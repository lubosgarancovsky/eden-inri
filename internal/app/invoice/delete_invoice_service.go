package invoice

import (
	"context"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

var _ ports.DeleteInvoiceUseCase = (*DeleteInvoiceService)(nil)

type DeleteInvoiceService struct {
	repo ports.PersistInvoicePort
	tm   ports.TransactionManager
}

func NewDeleteInvoiceService(repo ports.PersistInvoicePort, tm ports.TransactionManager) *DeleteInvoiceService {
	return &DeleteInvoiceService{repo: repo, tm: tm}
}

func (s *DeleteInvoiceService) Execute(ctx context.Context, userID, invoiceID uuid.UUID) error {
	return s.tm.WithTransaction(ctx, func(ctx context.Context) error { return s.repo.Delete(ctx, userID, invoiceID) })
}
