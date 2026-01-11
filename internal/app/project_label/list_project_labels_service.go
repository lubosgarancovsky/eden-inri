package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListProjectLabelsService struct {
	repo ports.PersistProjectLabelPort
}

func NewListProjectLabelsService(repo ports.PersistProjectLabelPort) *ListProjectLabelsService {
	return &ListProjectLabelsService{repo: repo}
}

func (s *ListProjectLabelsService) Execute(ctx context.Context, q *query.ListProjectScopedQuery) (*[]entity.ProjectLabel, int64, error) {
	return s.repo.List(ctx, q.ProjectID, q.ListingQuery)
}
