package kanban_board

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type UpdateKanbanBoardService struct {
	repo ports.PersistKanbanBoardPort
}

func NewUpdateKanbanBoardService(repo ports.PersistKanbanBoardPort) *UpdateKanbanBoardService {
	return &UpdateKanbanBoardService{repo: repo}
}

func (s *UpdateKanbanBoardService) Execute(ctx context.Context, cmd *command.UpdateKanbanBoardCommand) (*entity.KanbanBoard, error) {
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
