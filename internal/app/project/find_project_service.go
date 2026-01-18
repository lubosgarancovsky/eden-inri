package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindProjectByIDService struct {
	repo ports.PersistProjectPort
}

func NewFindProjectByIDService(repo ports.PersistProjectPort) *FindProjectByIDService {
	return &FindProjectByIDService{
		repo: repo,
	}
}

func (s *FindProjectByIDService) Execute(ctx context.Context, q *query.Query) (*entity.Project, error) {
	return s.repo.FindByID(ctx, q.UserID, q.ID)
}
