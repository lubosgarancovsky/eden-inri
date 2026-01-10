package ports

import (
	"context"

	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type CreateLabelUseCase interface {
	Execute(ctx context.Context, label *Label) (*Label, error)
}

type UpdateLabelUseCase interface {
	Execute(ctx context.Context, label *Label) (*Label, error)
}

type DeleteLabelUseCase interface {
	Execute(ctx context.Context, projectID, labelID uuid.UUID) error
}

type FindLabelByIDUseCase interface {
	Execute(ctx context.Context, projectID, labelID uuid.UUID) (*Label, error)
}

type ListLabelsUseCase interface {
	Execute(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]Label, int64, error)
}
