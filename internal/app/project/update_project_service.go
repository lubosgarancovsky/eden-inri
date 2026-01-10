package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.UpdateProjectUseCase = (*UpdateProjectService)(nil)

type UpdateProjectService struct {
	repository ports.PersistProjectPort
	tm         ports.TransactionManager
}

func NewUpdateProjectService(repository ports.PersistProjectPort, tm ports.TransactionManager) *UpdateProjectService {
	return &UpdateProjectService{repository: repository, tm: tm}
}

func (s *UpdateProjectService) Execute(ctx context.Context, cmd *command.UpdateProjectCommand) (*entity.Project, error) {
	project := cmd.ToDomain()
	if err := s.tm.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.Update(ctx, cmd.UserID, project); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return project, nil
}
