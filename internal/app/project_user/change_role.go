package project_user

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type ChangeProjectUserRoleService struct {
	repo        ports.PersistProjectUserPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewChangeProjectUserRoleService(repo ports.PersistProjectUserPort, hasRoleRepo ports.MemberHasRolePort) *ChangeProjectUserRoleService {
	return &ChangeProjectUserRoleService{repo: repo, hasRoleRepo: hasRoleRepo}
}

func (s *ChangeProjectUserRoleService) Execute(ctx context.Context, cmd *command.ChangeProjectUserRoleCommand) (*entity.ProjectUser, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return nil, err
	}

	if !hasRole {
		return nil, app_err.ErrInsufficientProjectRole
	}

	projectUser, err := s.repo.FindByID(ctx, cmd.ProjectID, cmd.MemberID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(projectUser)
	if err = s.repo.SetRole(ctx, cmd.ProjectID, cmd.MemberID, projectUser.Role); err != nil {
		return nil, err
	}

	return projectUser, nil
}
