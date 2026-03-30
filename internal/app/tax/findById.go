package tax

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.FindTaxByIDUseCase = (*FindTaxByIDService)(nil)

type FindTaxByIDService struct {
	repo ports.TaxRepository
}

func NewFindTaxByIDService(repo ports.TaxRepository) *FindTaxByIDService {
	return &FindTaxByIDService{repo: repo}
}

func (s *FindTaxByIDService) Execute(ctx context.Context, query *query.Query) (*entity.Tax, error) {
	return s.repo.FindByID(ctx, query.UserID, query.ID)
}
