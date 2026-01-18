package invoice_stats

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type GetInvoiceStatsService struct {
	repo ports.PersistInvoiceStatsPort
}

func NewGetInvoiceStatsService(repo ports.PersistInvoiceStatsPort) *GetInvoiceStatsService {
	return &GetInvoiceStatsService{repo: repo}
}

func (s *GetInvoiceStatsService) Execute(ctx context.Context, query *query.ListQuery) (*entity.InvoiceStats, error) {
	return s.repo.Stats(ctx, query.UserID, query.ListingQuery)
}
