package tax

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.ListTaxesUseCase = (*ListTaxesService)(nil)

type ListTaxesService struct {
	repo ports.TaxRepository
}

func NewListTaxesService(repo ports.TaxRepository) *ListTaxesService {
	return &ListTaxesService{repo: repo}
}

func (s *ListTaxesService) Execute(ctx context.Context, query *query.ListQuery) (*[]entity.Tax, int64, error) {
	return s.repo.List(ctx, query.UserID, query.ListingQuery)
}
