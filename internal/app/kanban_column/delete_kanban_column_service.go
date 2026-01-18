package kanban_column

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type DeleteKanbanColumnService struct {
	repo        ports.PersistKanbanColumnPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewDeleteKanbanColumnService(repo ports.PersistKanbanColumnPort, hasRoleRepo ports.MemberHasRolePort) *DeleteKanbanColumnService {
	return &DeleteKanbanColumnService{repo: repo, hasRoleRepo: hasRoleRepo}
}

func (s *DeleteKanbanColumnService) Execute(ctx context.Context, cmd *command.DeleteKanbanColumnCommand) error {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return err
	}

	if !hasRole {
		return app_err.ErrNotAMember
	}

	return s.repo.Delete(ctx, cmd.BoardID, cmd.ID)
}
