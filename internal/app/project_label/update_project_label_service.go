package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type UpdateProjectLabelService struct {
	repo ports.PersistProjectLabelPort
}

func NewUpdateProjectLabelService(repo ports.PersistProjectLabelPort) *UpdateProjectLabelService {
	return &UpdateProjectLabelService{repo: repo}
}

func (s *UpdateProjectLabelService) Execute(ctx context.Context, cmd *command.UpdateProjectLabelCommand) (*entity.ProjectLabel, error) {
	label, err := s.repo.FindByID(ctx, cmd.ProjectID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(label)

	if err := s.repo.Update(ctx, label); err != nil {
		return nil, err
	}

	return label, nil
}
