package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

var _ ports.DeleteProjectUseCase = (*DeleteProjectService)(nil)

type DeleteProjectService struct {
	repository ports.PersistProjectPort
}

func NewDeleteProjectService(repository ports.PersistProjectPort) *DeleteProjectService {
	return &DeleteProjectService{repository: repository}
}

func (s *DeleteProjectService) Execute(ctx context.Context, cmd *command.DeleteProjectCommand) error {
	return s.repository.Delete(ctx, cmd.UserID, cmd.ProjectID)
}
