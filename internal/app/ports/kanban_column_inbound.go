package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateKanbanColumnUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateKanbanColumnCommand) (*entity.KanbanColumn, error)
}

type UpdateKanbanColumnUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateKanbanColumnCommand) (*entity.KanbanColumn, error)
}

type DeleteKanbanColumnUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteBoardScopedCommand) error
}

type FindKanbanColumnByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindByIDBoardScopedQuery) (*entity.KanbanColumn, error)
}

type ListKanbanColumnsUseCase interface {
	Execute(ctx context.Context, boardID uuid.UUID) ([]entity.KanbanColumn, error)
}
