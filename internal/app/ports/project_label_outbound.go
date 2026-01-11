package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistProjectLabelPort interface {
	Create(ctx context.Context, label *entity.ProjectLabel) error
	Update(ctx context.Context, label *entity.ProjectLabel) error
	Delete(ctx context.Context, projectID, labelID uuid.UUID) error
	FindByID(ctx context.Context, projectID, labelID uuid.UUID) (*entity.ProjectLabel, error)
	List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ProjectLabel, int64, error)
}
