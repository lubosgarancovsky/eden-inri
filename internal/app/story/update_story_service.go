package story

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type UpdateStoryService struct {
	repo       ports.PersistStoryPort
	memberRepo ports.MemberHasRolePort
}

func NewUpdateStoryService(repo ports.PersistStoryPort, memberRepo ports.MemberHasRolePort) *UpdateStoryService {
	return &UpdateStoryService{repo: repo, memberRepo: memberRepo}
}

func (s *UpdateStoryService) Execute(ctx context.Context, cmd *command.UpdateStoryCommand) (*entity.Story, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin, entity.ProjectRoleDeveloper}
	hasRole, err := s.memberRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return nil, err
	}
	if !hasRole {
		return nil, app_err.ErrInsufficientProjectRole
	}

	story, err := s.repo.FindByID(ctx, cmd.ProjectID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(story)

	if err := s.repo.Update(ctx, story); err != nil {
		return nil, err
	}

	return story, nil
}
