package kanban_board

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type CreateKanbanBoardService struct {
	repo        ports.PersistKanbanBoardPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewCreateKanbanBoardService(repo ports.PersistKanbanBoardPort, hasRoleRepo ports.MemberHasRolePort) *CreateKanbanBoardService {
	return &CreateKanbanBoardService{repo: repo, hasRoleRepo: hasRoleRepo}
}

func (s *CreateKanbanBoardService) Execute(ctx context.Context, cmd *command.CreateKanbanBoardCommand) (*entity.KanbanBoard, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return nil, err
	}

	if !hasRole {
		return nil, app_err.ErrInsufficientProjectRole
	}

	board := cmd.ToDomain()
	if err := s.repo.Create(ctx, board); err != nil {
		return nil, err
	}
	return board, nil
}
