package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type UpdateProjectService struct {
	repo ports.PersistProjectPort
}

func NewUpdateProjectService(repo ports.PersistProjectPort) *UpdateProjectService {
	return &UpdateProjectService{
		repo: repo,
	}
}

func (s *UpdateProjectService) Execute(ctx context.Context, cmd *command.UpdateProjectCommand) (*entity.Project, error) {
	project, err := s.repo.FindByID(ctx, cmd.UserID, cmd.ID)
	if err != nil {
		return nil, err
	}

	if !project.CanMutate() {
		return nil, go_kit.ErrForbidden.WithMessage("you don't have sufficient permission to mutate this project")
	}

	cmd.Apply(project)

	if err := s.repo.Update(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}
