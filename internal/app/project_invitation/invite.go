package project_invitation

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type InviteUserToProjectService struct {
	repo        ports.PersistProjectInvitationPort
	userRepo    ports.QueryUserPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewInviteUserToProjectService(repo ports.PersistProjectInvitationPort, userRepo ports.QueryUserPort, hasRoleRepo ports.MemberHasRolePort) *InviteUserToProjectService {
	return &InviteUserToProjectService{repo: repo, userRepo: userRepo, hasRoleRepo: hasRoleRepo}
}

func (s *InviteUserToProjectService) Execute(ctx context.Context, cmd *command.InviteUserToProjectCommand) error {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return err
	}

	if !hasRole {
		return app_err.ErrInsufficientProjectRole
	}

	user, err := s.userRepo.FindByEmail(ctx, cmd.Email)
	if err != nil {
		// Do not return error when user does not exist to prevent brute force email checking
		return nil
	}

	invitation := entity.NewProjectInvitation(cmd.ProjectID, user.ID, cmd.ProjectID, cmd.Role)
	if err := s.repo.Save(ctx, invitation); err != nil {
		return err
	}

	// TODO: Create event for email service to handle

	return nil
}
