package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistProjectPort interface {
	Create(ctx context.Context, project *entity.Project) error
	Update(ctx context.Context, project *entity.Project) error
	Delete(ctx context.Context, projectID uuid.UUID) error
	FindByID(ctx context.Context, userID, projectID uuid.UUID) (*entity.Project, error)
	List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Project, int64, error)
}
