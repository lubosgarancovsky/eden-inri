package invoice

import (
	"context"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

var _ ports.FindInvoiceByIDUseCase = (*FindInvoiceService)(nil)

type FindInvoiceService struct{ repo ports.QueryInvoicePort }

func NewFindInvoiceService(repo ports.QueryInvoicePort) *FindInvoiceService {
	return &FindInvoiceService{repo}
}

func (s *FindInvoiceService) Execute(ctx context.Context, userID, invoiceID uuid.UUID) (*ports.Invoice, error) {
	return s.repo.FindByID(ctx, userID, invoiceID)
}
