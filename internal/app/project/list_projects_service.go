package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.ListProjectsUseCase = (*ListProjectsService)(nil)

type ListProjectsService struct {
	repository ports.QueryProjectPort
}

func NewListProjectsService(repository ports.QueryProjectPort) *ListProjectsService {
	return &ListProjectsService{repository}
}

func (s *ListProjectsService) Execute(ctx context.Context, q *query.ListProjectsQuery) (*[]entity.Project, int64, error) {
	lq := q.ListingQuery // copy value to get pointer below
	return s.repository.List(ctx, q.UserID, &lq)
}
