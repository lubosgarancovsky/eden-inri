package kanban_column

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type UpdateKanbanColumnService struct {
	repo        ports.PersistKanbanColumnPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewUpdateKanbanColumnService(repo ports.PersistKanbanColumnPort, hasRoleRepo ports.MemberHasRolePort) *UpdateKanbanColumnService {
	return &UpdateKanbanColumnService{repo: repo, hasRoleRepo: hasRoleRepo}
}

func (s *UpdateKanbanColumnService) Execute(ctx context.Context, cmd *command.UpdateKanbanColumnCommand) (*entity.KanbanColumn, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return nil, err
	}

	if !hasRole {
		return nil, app_err.ErrNotAMember
	}

	col, err := s.repo.FindByID(ctx, cmd.BoardID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(col)

	if err := s.repo.Update(ctx, col); err != nil {
		return nil, err
	}

	return col, nil
}
