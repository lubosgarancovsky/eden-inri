package story

import (
	"context"
	"fmt"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type ChangeStoryAssigneeService struct {
	repo       ports.PersistStoryPort
	memberRepo ports.MemberHasRolePort
}

func NewChangeStoryAssigneeService(repo ports.PersistStoryPort, memberRepo ports.MemberHasRolePort) *ChangeStoryAssigneeService {
	return &ChangeStoryAssigneeService{repo: repo, memberRepo: memberRepo}
}

func (s *ChangeStoryAssigneeService) Execute(ctx context.Context, cmd *command.ChangeStoryAssigneeCommand) error {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin, entity.ProjectRoleDeveloper}
	hasRole, err := s.memberRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return err
	}
	if !hasRole {
		return fmt.Errorf("user is not a member of project")
	}

	return s.repo.ChangeAssignee(ctx, cmd.ProjectID, cmd.ID, cmd.AssigneeID)
}
