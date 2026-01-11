package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.CreateProjectLabelUseCase = (*CreateProjectLabelService)(nil)

type CreateProjectLabelService struct {
	repo ports.PersistLabelPort
}

func NewCreateProjectLabelService(repo ports.PersistLabelPort) *CreateProjectLabelService {
	return &CreateProjectLabelService{repo: repo}
}

func (s *CreateProjectLabelService) Execute(ctx context.Context, cmd *command.CreateProjectLabelCommand) (*entity.ProjectLabel, error) {
	label := cmd.ToDomain()
	if err := s.repo.Create(ctx, label); err != nil {
		return nil, err
	}
	return label, nil
}
