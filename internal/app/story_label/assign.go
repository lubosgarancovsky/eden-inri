package story_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type AssignStoryLabelService struct {
	repo        ports.PersistStoryLabelPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewAssignStoryLabelService(repo ports.PersistStoryLabelPort, hasRoleRepo ports.MemberHasRolePort) *AssignStoryLabelService {
	return &AssignStoryLabelService{repo, hasRoleRepo}
}

func (s *AssignStoryLabelService) Execute(ctx context.Context, cmd *command.StoryLabelCommand) error {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin, entity.ProjectRoleDeveloper}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return err
	}
	if !hasRole {
		return app_error.ErrInsufficientProjectRole
	}

	return s.repo.Assign(ctx, cmd.StoryID, cmd.LabelID)
}
