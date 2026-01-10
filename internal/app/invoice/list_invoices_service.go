package invoice

import (
	"context"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

var _ ports.ListInvoicesUseCase = (*ListInvoicesService)(nil)

type ListInvoicesService struct{ repo ports.QueryInvoicePort }

func NewListInvoicesService(repo ports.QueryInvoicePort) *ListInvoicesService {
	return &ListInvoicesService{repo}
}

func (s *ListInvoicesService) Execute(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]ports.Invoice, int64, error) {
	return s.repo.List(ctx, userID, lq)
}
