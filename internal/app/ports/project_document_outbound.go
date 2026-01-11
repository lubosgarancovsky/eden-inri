package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistProjectDocumentPort interface {
	Create(ctx context.Context, doc *entity.ProjectDocument) error
	Update(ctx context.Context, doc *entity.ProjectDocument) error
	Delete(ctx context.Context, projectID, documentID uuid.UUID) error
	FindByID(ctx context.Context, projectID, documentID uuid.UUID) (*entity.ProjectDocument, error)
	List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ProjectDocument, int64, error)
}
