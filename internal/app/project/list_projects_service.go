package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListProjectsService struct {
	repo ports.PersistProjectPort
}

func NewListProjectsService(repo ports.PersistProjectPort) *ListProjectsService {
	return &ListProjectsService{
		repo: repo,
	}
}

func (s *ListProjectsService) Execute(ctx context.Context, q *query.ListQuery) (*[]entity.Project, int64, error) {
	return s.repo.List(ctx, q.UserID, q.ListingQuery)
}
