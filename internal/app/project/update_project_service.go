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
}

func NewUpdateProjectService(repository ports.PersistProjectPort) *UpdateProjectService {
	return &UpdateProjectService{repository: repository}
}

func (s *UpdateProjectService) Execute(ctx context.Context, cmd *command.UpdateProjectCommand) (*entity.Project, error) {
	project, err := s.repository.FindByID(ctx, cmd.UserID, cmd.ProjectID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(project)

	if err = s.repository.Update(ctx, cmd.UserID, project); err != nil {
		return nil, err
	}

	return project, nil
}
