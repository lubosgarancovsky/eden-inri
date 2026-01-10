package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

var _ ports.DeleteProjectUseCase = (*DeleteProjectService)(nil)

type DeleteProjectService struct {
	repository ports.PersistProjectPort
	tm         ports.TransactionManager
}

func NewDeleteProjectService(repository ports.PersistProjectPort, tm ports.TransactionManager) *DeleteProjectService {
	return &DeleteProjectService{repository: repository, tm: tm}
}

func (s *DeleteProjectService) Execute(ctx context.Context, cmd *command.DeleteProjectCommand) error {
	return s.tm.WithTransaction(ctx, func(ctx context.Context) error {
		return s.repository.Delete(ctx, cmd.UserID, cmd.ProjectID)
	})
}
