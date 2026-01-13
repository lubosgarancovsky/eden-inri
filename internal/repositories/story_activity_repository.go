package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StoryActivityRepository struct {
	db *gorm.DB
}

func NewStoryActivityRepository(db *gorm.DB) *StoryActivityRepository {
	return &StoryActivityRepository{db: db}
}

func (r *StoryActivityRepository) Insert(ctx context.Context, activity *models.StoryActivity) (*models.StoryActivity, error) {
	err := r.db.
		WithContext(ctx).
		Create(activity).
		Error

	return activity, err
}

func (r *StoryActivityRepository) Update(ctx context.Context, activity *models.StoryActivity) (*models.StoryActivity, error) {
	err := r.db.
		WithContext(ctx).
		Model(activity).
		Clauses(clause.Returning{}).
		Where("story_id = ? AND actor_id = ? AND id = ?", activity.StoryID, activity.ActorID, activity.ID).
		Updates(activity).
		Error

	return activity, err
}

func (r *StoryActivityRepository) FindByID(ctx context.Context, userID, storyID, activityID uuid.UUID) (*models.StoryActivity, error) {
	var activity models.StoryActivity
	err := r.db.
		WithContext(ctx).
		Model(&activity).
		Where("story_id = ? AND actor_id = ? AND id = ?", storyID, userID, activityID).
		Find(&activity).
		Error

	return &activity, err
}

func (r *StoryActivityRepository) ListActivities(ctx context.Context, storyID uuid.UUID, lq *list.ListingQuery) ([]models.StoryActivity, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(models.StoryActivity{}).
		Preload("Actor").
		Where("story_id = ?", storyID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := postgres.List[models.StoryActivity](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
