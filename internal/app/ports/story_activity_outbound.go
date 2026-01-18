package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
)

type PersistStoryActivityPort interface {
	Create(ctx context.Context, activity *entity.StoryActivity) error
	Update(ctx context.Context, activity *entity.StoryActivity) error
	Delete(ctx context.Context, activityID uuid.UUID) error
	FindByID(ctx context.Context, activityID uuid.UUID) (*entity.StoryActivity, error)
	List(ctx context.Context, storyID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.StoryActivity, int64, error)
}
