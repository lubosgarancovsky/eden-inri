package kanban_board

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type DeleteKanbanBoardService struct {
	repo        ports.PersistKanbanBoardPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewDeleteKanbanBoardService(repo ports.PersistKanbanBoardPort, hasRoleRepo ports.MemberHasRolePort) *DeleteKanbanBoardService {
	return &DeleteKanbanBoardService{repo: repo}
}

func (s *DeleteKanbanBoardService) Execute(ctx context.Context, cmd *command.DeleteKanbanBoardCommand) error {
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
