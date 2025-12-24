package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
)

type StoryActivityRepository struct {
	db *gorm.DB
}

func NewStoryActivityRepository(db *gorm.DB) *StoryActivityRepository {
	return &StoryActivityRepository{db: db}
}

func (r *StoryActivityRepository) Insert(ctx context.Context, activity *model.StoryActivity) (*model.StoryActivity, error) {
	err := r.db.
		WithContext(ctx).
		Create(activity).
		Error

	return activity, err
}

func (r *StoryActivityRepository) ListActivities(ctx context.Context, storyID uuid.UUID, lq *list.ListingQuery) ([]model.StoryActivity, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(model.StoryActivity{}).
		Where("story_id = ?", storyID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.StoryActivity](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
