package ports

import (
	"context"

	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistLabelPort interface {
	Create(ctx context.Context, label *Label) error
	Update(ctx context.Context, label *Label) error
	Delete(ctx context.Context, projectID, labelID uuid.UUID) error
}

type QueryLabelPort interface {
	FindByID(ctx context.Context, projectID, labelID uuid.UUID) (*Label, error)
	List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]Label, int64, error)
}

// Lightweight domain for Label to decouple from legacy models
type Label struct {
	ID          uuid.UUID
	ProjectID   uuid.UUID
	Name        string
	Description string
	Color       string
}
