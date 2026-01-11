package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateProjectLabelService struct {
	repo ports.PersistProjectLabelPort
}

func NewCreateProjectLabelService(repo ports.PersistProjectLabelPort) *CreateProjectLabelService {
	return &CreateProjectLabelService{repo: repo}
}

func (s *CreateProjectLabelService) Execute(ctx context.Context, cmd *command.CreateProjectLabelCommand) (*entity.ProjectLabel, error) {
	label := cmd.ToDomain()
	if err := s.repo.Create(ctx, label); err != nil {
		return nil, err
	}
	return label, nil
}
