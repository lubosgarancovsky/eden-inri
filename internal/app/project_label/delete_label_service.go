package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

var _ ports.DeleteProjectLabelUseCase = (*DeleteProjectLabelService)(nil)

type DeleteProjectLabelService struct {
	repo ports.PersistLabelPort
}

func NewDeleteProjectLabelService(repo ports.PersistLabelPort) *DeleteProjectLabelService {
	return &DeleteProjectLabelService{repo: repo}
}

func (s *DeleteProjectLabelService) Execute(ctx context.Context, cmd *command.DeleteProjectLabelCommand) error {
	return s.repo.Delete(ctx, cmd.ProjectID, cmd.ID)
}
