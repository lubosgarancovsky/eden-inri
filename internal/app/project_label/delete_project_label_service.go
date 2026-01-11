package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

type DeleteProjectLabelService struct {
	repo ports.PersistProjectLabelPort
}

func NewDeleteProjectLabelService(repo ports.PersistProjectLabelPort) *DeleteProjectLabelService {
	return &DeleteProjectLabelService{repo: repo}
}

func (s *DeleteProjectLabelService) Execute(ctx context.Context, cmd *command.DeleteProjectScopedCommand) error {
	return s.repo.Delete(ctx, cmd.ProjectID, cmd.ID)
}
