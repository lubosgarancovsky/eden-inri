package tax_stats

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type GetTaxStatsService struct {
	repo ports.PersistTaxStatsPort
}

func NewGetTaxStatsService(repo ports.PersistTaxStatsPort) *GetTaxStatsService {
	return &GetTaxStatsService{repo: repo}
}

func (s *GetTaxStatsService) Execute(ctx context.Context, query *query.ListQuery) (*entity.TaxStats, error) {
	return s.repo.Stats(ctx, query.UserID, query.ListingQuery)
}
