package kanban_board

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

type DeleteKanbanBoardService struct {
	repo ports.PersistKanbanBoardPort
}

func NewDeleteKanbanBoardService(repo ports.PersistKanbanBoardPort) *DeleteKanbanBoardService {
	return &DeleteKanbanBoardService{repo: repo}
}

func (s *DeleteKanbanBoardService) Execute(ctx context.Context, cmd *command.DeleteProjectScopedCommand) error {
	return s.repo.Delete(ctx, cmd.ProjectID, cmd.ID)
}
