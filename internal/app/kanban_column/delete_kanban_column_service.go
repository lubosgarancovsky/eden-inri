package kanban_column

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

type DeleteKanbanColumnService struct {
	repo ports.PersistKanbanColumnPort
}

func NewDeleteKanbanColumnService(repo ports.PersistKanbanColumnPort) *DeleteKanbanColumnService {
	return &DeleteKanbanColumnService{repo: repo}
}

func (s *DeleteKanbanColumnService) Execute(ctx context.Context, cmd *command.DeleteBoardScopedCommand) error {
	return s.repo.Delete(ctx, cmd.BoardID, cmd.ID)
}
