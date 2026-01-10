package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistClientPort interface {
	Create(ctx context.Context, client *entity.Client) error
	Update(ctx context.Context, client *entity.Client) error
	Delete(ctx context.Context, userID, clientID uuid.UUID) error
}

type QueryClientPort interface {
	FindByID(ctx context.Context, userID, clientID uuid.UUID) (*entity.Client, error)
	List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Client, int64, error)
}
