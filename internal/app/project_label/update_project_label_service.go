package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type UpdateProjectLabelService struct {
	repo        ports.PersistProjectLabelPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewUpdateProjectLabelService(repo ports.PersistProjectLabelPort, hasRoleRepo ports.MemberHasRolePort) *UpdateProjectLabelService {
	return &UpdateProjectLabelService{repo: repo, hasRoleRepo: hasRoleRepo}
}

func (s *UpdateProjectLabelService) Execute(ctx context.Context, cmd *command.UpdateProjectLabelCommand) (*entity.ProjectLabel, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return nil, err
	}

	if !hasRole {
		return nil, app_err.ErrInsufficientProjectRole
	}

	label, err := s.repo.FindByID(ctx, cmd.ProjectID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(label)

	if err = s.repo.Update(ctx, label); err != nil {
		return nil, err
	}

	return label, nil
}
