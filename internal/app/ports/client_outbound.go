package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type PersistClientPort interface {
	Create(ctx context.Context, client *entity.Client) error
	Update(ctx context.Context, client *entity.Client) error
	Delete(ctx context.Context, userID, clientID uuid.UUID) error
}

type QueryClientPort interface {
	FindByID(ctx context.Context, clientID uuid.UUID) (*entity.Client, error)
	List(ctx context.Context, offset, limit int) ([]entity.Client, int, error)
}
