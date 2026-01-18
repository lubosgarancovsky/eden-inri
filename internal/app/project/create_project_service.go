package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateProjectService struct {
	projectRepo     ports.PersistProjectPort
	projectUserRepo ports.PersistProjectUserPort
	tx              ports.TransactionManager
}

func NewCreateProjectService(projectRepo ports.PersistProjectPort, projectUserRepo ports.PersistProjectUserPort, tx ports.TransactionManager) *CreateProjectService {
	return &CreateProjectService{
		projectRepo:     projectRepo,
		projectUserRepo: projectUserRepo,
		tx:              tx,
	}
}

func (s *CreateProjectService) Execute(ctx context.Context, cmd *command.CreateProjectCommand) (*entity.Project, error) {
	project := cmd.ToDomain()

	err := s.tx.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := s.projectRepo.Create(txCtx, project); err != nil {
			return err
		}

		projectUser := entity.NewProjectOwner(cmd.UserID, project.ID)
		if err := s.projectUserRepo.Insert(txCtx, projectUser); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return project, nil
}
