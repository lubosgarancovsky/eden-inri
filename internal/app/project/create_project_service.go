package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.CreateProjectUseCase = (*CreateProjectService)(nil)

type CreateProjectService struct {
	repository      ports.PersistProjectPort
	projectUserRepo ports.ProjectUserPersistPort
	tm              ports.TransactionManager
}

func NewCreateProjectService(repository ports.PersistProjectPort, projectUserRepo ports.ProjectUserPersistPort, tm ports.TransactionManager) *CreateProjectService {
	return &CreateProjectService{repository: repository, projectUserRepo: projectUserRepo, tm: tm}
}

func (s *CreateProjectService) Execute(ctx context.Context, cmd *command.CreateProjectCommand) (*entity.Project, error) {
	project := cmd.ToDomain()
	if err := s.tm.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.Create(ctx, project); err != nil {
			return err
		}
		if err := s.projectUserRepo.AddOwner(ctx, cmd.UserID, project.ID); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return project, nil
}
