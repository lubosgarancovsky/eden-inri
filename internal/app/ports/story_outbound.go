package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
)

type PersistStoryPort interface {
	Create(ctx context.Context, story *entity.Story) error
	Update(ctx context.Context, story *entity.Story) error
	Delete(ctx context.Context, projectID, storyID uuid.UUID) error
	FindByID(ctx context.Context, projectID, storyID uuid.UUID) (*entity.Story, error)
	FindBySlug(ctx context.Context, projectID uuid.UUID, slug string) (*entity.Story, error)
	List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Story, int64, error)
	ListAssigned(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Story, int64, error)
	ChangeAssignee(ctx context.Context, projectID, storyID uuid.UUID, assigneeID *uuid.UUID) error
}
