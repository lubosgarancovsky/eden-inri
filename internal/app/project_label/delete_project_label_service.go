package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type DeleteProjectLabelService struct {
	repo        ports.PersistProjectLabelPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewDeleteProjectLabelService(repo ports.PersistProjectLabelPort, hasRoleRepo ports.MemberHasRolePort) *DeleteProjectLabelService {
	return &DeleteProjectLabelService{repo, hasRoleRepo}
}

func (s *DeleteProjectLabelService) Execute(ctx context.Context, cmd *command.DeleteProjectLabelCommand) error {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return err
	}

	if !hasRole {
		return app_err.ErrInsufficientProjectRole
	}

	return s.repo.Delete(ctx, cmd.ProjectID, cmd.ID)
}
