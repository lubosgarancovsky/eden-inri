package business_entity

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.ListBusinessEntitiesUseCase = (*ListBusinessEntitiesService)(nil)

type ListBusinessEntitiesService struct {
	repo ports.BusinessEntityRepository
}

func NewListBusinessEntitiesService(repo ports.BusinessEntityRepository) *ListBusinessEntitiesService {
	return &ListBusinessEntitiesService{repo: repo}
}

func (s *ListBusinessEntitiesService) Execute(ctx context.Context, query *query.ListQuery) (*[]entity.BusinessEntity, int64, error) {
	return s.repo.List(ctx, query.UserID, query.ListingQuery)
}
