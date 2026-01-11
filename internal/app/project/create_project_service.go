package project

import (
	"context"
	"time"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateProjectService struct {
	repo ports.PersistProjectPort
	tx   ports.TransactionManager
}

func NewCreateProjectService(repo ports.PersistProjectPort, tx ports.TransactionManager) *CreateProjectService {
	return &CreateProjectService{
		repo: repo,
		tx:   tx,
	}
}

func (s *CreateProjectService) Execute(ctx context.Context, cmd *command.CreateProjectCommand) (*entity.Project, error) {
	project := cmd.ToDomain()

	err := s.tx.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := s.repo.Create(txCtx, project); err != nil {
			return err
		}

		pu := &entity.ProjectUser{
			ProjectID: project.ID,
			UserID:    cmd.UserID,
			IsStarred: false,
			Role:      entity.ProjectRoleOwner,
			JoinedAt:  time.Now(),
		}

		if err := s.repo.InsertProjectUser(txCtx, pu); err != nil {
			return err
		}

		project.Role = pu.Role
		project.IsStarred = pu.IsStarred

		return nil
	})

	if err != nil {
		return nil, err
	}

	return project, nil
}
