package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindProjectLabelByIDService struct {
	repo ports.PersistProjectLabelPort
}

func NewFindProjectLabelByIDService(repo ports.PersistProjectLabelPort) *FindProjectLabelByIDService {
	return &FindProjectLabelByIDService{repo: repo}
}

func (s *FindProjectLabelByIDService) Execute(ctx context.Context, q *query.FindByIDProjectScopedQuery) (*entity.ProjectLabel, error) {
	return s.repo.FindByID(ctx, q.ProjectID, q.ID)
}
