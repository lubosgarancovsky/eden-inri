package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
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
	// Note: In hexagonal, the service should handle business logic.
	// We might need to check if the project exists or if the user has permissions.
	// For now, let's keep it simple as in legacy.
	project := &entity.Project{
		ID: cmd.ID,
	}
	cmd.Apply(project)

	if err := s.repo.Update(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}
