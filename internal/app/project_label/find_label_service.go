package project_label

import (
	"context"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

var _ ports.FindLabelByIDUseCase = (*FindLabelService)(nil)

type FindLabelService struct{ repo ports.QueryLabelPort }

func NewFindLabelService(repo ports.QueryLabelPort) *FindLabelService { return &FindLabelService{repo} }

func (s *FindLabelService) Execute(ctx context.Context, projectID, labelID uuid.UUID) (*ports.Label, error) {
	return s.repo.FindByID(ctx, projectID, labelID)
}
