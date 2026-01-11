package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistKanbanBoardPort interface {
	Create(ctx context.Context, board *entity.KanbanBoard) error
	Update(ctx context.Context, board *entity.KanbanBoard) error
	Delete(ctx context.Context, projectID, boardID uuid.UUID) error
	FindByID(ctx context.Context, projectID, boardID uuid.UUID) (*entity.KanbanBoard, error)
	List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.KanbanBoard, int64, error)
}
