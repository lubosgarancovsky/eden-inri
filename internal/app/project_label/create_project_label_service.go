package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type CreateProjectLabelService struct {
	repo        ports.PersistProjectLabelPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewCreateProjectLabelService(repo ports.PersistProjectLabelPort, hasRoleRepo ports.MemberHasRolePort) *CreateProjectLabelService {
	return &CreateProjectLabelService{repo: repo, hasRoleRepo: hasRoleRepo}
}

func (s *CreateProjectLabelService) Execute(ctx context.Context, cmd *command.CreateProjectLabelCommand) (*entity.ProjectLabel, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return nil, err
	}

	if !hasRole {
		return nil, app_err.ErrInsufficientProjectRole
	}

	label := cmd.ToDomain()
	if err := s.repo.Create(ctx, label); err != nil {
		return nil, err
	}
	return label, nil
}
