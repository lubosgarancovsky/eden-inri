package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateKanbanBoardUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateKanbanBoardCommand) (*entity.KanbanBoard, error)
}

type UpdateKanbanBoardUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateKanbanBoardCommand) (*entity.KanbanBoard, error)
}

type DeleteKanbanBoardUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteKanbanBoardCommand) error
}

type FindKanbanBoardByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindByIDKanbanBoardQuery) (*entity.KanbanBoard, error)
}

type ListKanbanBoardsUseCase interface {
	Execute(ctx context.Context, query *query.ListKanbanBoardQuery) (*[]entity.KanbanBoard, int64, error)
}
