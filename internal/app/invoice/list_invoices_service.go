package invoice

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.ListInvoicesUseCase = (*ListInvoicesService)(nil)

type ListInvoicesService struct{ repo ports.QueryInvoicePort }

func NewListInvoicesService(repo ports.QueryInvoicePort) *ListInvoicesService {
	return &ListInvoicesService{repo}
}

func (s *ListInvoicesService) Execute(ctx context.Context, query *query.ListQuery) (*[]entity.Invoice, int64, error) {
	return s.repo.List(ctx, query.UserID, query.ListingQuery)
}
