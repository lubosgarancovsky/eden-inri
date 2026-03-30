package business_entity

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.FindBusinessEntityByIDUseCase = (*FindBusinessEntityByIDService)(nil)

type FindBusinessEntityByIDService struct {
	repo ports.BusinessEntityRepository
}

func NewFindBusinessEntityByIDService(repo ports.BusinessEntityRepository) *FindBusinessEntityByIDService {
	return &FindBusinessEntityByIDService{repo: repo}
}

func (s *FindBusinessEntityByIDService) Execute(ctx context.Context, query *query.Query) (*entity.BusinessEntity, error) {
	return s.repo.FindByID(ctx, query.UserID, query.ID)
}
