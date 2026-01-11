package kanban_board

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateKanbanBoardService struct {
	repo ports.PersistKanbanBoardPort
}

func NewCreateKanbanBoardService(repo ports.PersistKanbanBoardPort) *CreateKanbanBoardService {
	return &CreateKanbanBoardService{repo: repo}
}

func (s *CreateKanbanBoardService) Execute(ctx context.Context, cmd *command.CreateKanbanBoardCommand) (*entity.KanbanBoard, error) {
	board := cmd.ToDomain()
	if err := s.repo.Create(ctx, board); err != nil {
		return nil, err
	}
	return board, nil
}
