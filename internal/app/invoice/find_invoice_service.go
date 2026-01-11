package invoice

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.FindInvoiceByIDUseCase = (*FindInvoiceService)(nil)

type FindInvoiceService struct{ repo ports.PersistInvoicePort }

func NewFindInvoiceService(repo ports.PersistInvoicePort) *FindInvoiceService {
	return &FindInvoiceService{repo}
}

func (s *FindInvoiceService) Execute(ctx context.Context, query *query.FindByIDQuery) (*entity.Invoice, error) {
	return s.repo.FindByID(ctx, query.UserID, query.ID)
}
