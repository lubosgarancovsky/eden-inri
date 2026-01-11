package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type PersistKanbanColumnPort interface {
	Create(ctx context.Context, col *entity.KanbanColumn) error
	Update(ctx context.Context, col *entity.KanbanColumn) error
	Delete(ctx context.Context, boardID, columnID uuid.UUID) error
	FindByID(ctx context.Context, boardID, columnID uuid.UUID) (*entity.KanbanColumn, error)
	List(ctx context.Context, boardID uuid.UUID) ([]entity.KanbanColumn, error)
}
