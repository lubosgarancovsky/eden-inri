package kanban_board

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type UpdateKanbanBoardService struct {
	repo        ports.PersistKanbanBoardPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewUpdateKanbanBoardService(repo ports.PersistKanbanBoardPort, hasRoleRepo ports.MemberHasRolePort) *UpdateKanbanBoardService {
	return &UpdateKanbanBoardService{repo: repo, hasRoleRepo: hasRoleRepo}
}

func (s *UpdateKanbanBoardService) Execute(ctx context.Context, cmd *command.UpdateKanbanBoardCommand) (*entity.KanbanBoard, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return nil, err
	}

	if !hasRole {
		return nil, app_err.ErrInsufficientProjectRole
	}

	board, err := s.repo.FindByID(ctx, cmd.ProjectID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(board)

	if err := s.repo.Update(ctx, board); err != nil {
		return nil, err
	}

	return board, nil
}
