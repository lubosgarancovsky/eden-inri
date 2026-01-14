package project_user

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type DeleteProjectUserService struct {
	repo        ports.PersistProjectUserPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewDeleteProjectUserService(repo ports.PersistProjectUserPort, hasRoleRepo ports.MemberHasRolePort) *DeleteProjectUserService {
	return &DeleteProjectUserService{repo: repo, hasRoleRepo: hasRoleRepo}
}

func (s *DeleteProjectUserService) Execute(ctx context.Context, cmd *command.ScopedCommand) error {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ScopeID, roles)
	if err != nil {
		return err
	}

	if !hasRole {
		return app_err.ErrInsufficientProjectRole
	}

	// ctx, ProjectID, MemberID
	return s.repo.Delete(ctx, cmd.ScopeID, cmd.ID)
}
