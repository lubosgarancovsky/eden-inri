package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.FindProjectByIDUseCase = (*FindProjectByIDService)(nil)

type FindProjectByIDService struct {
	repository ports.QueryProjectPort
}

func NewFindProjectByIDService(repository ports.QueryProjectPort) *FindProjectByIDService {
	return &FindProjectByIDService{repository}
}

func (s *FindProjectByIDService) Execute(ctx context.Context, q *query.FindProjectByIDQuery) (*entity.Project, error) {
	return s.repository.FindByID(ctx, q.UserID, q.ProjectID)
}
