package repository

import (
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

// Insert a new activity
func (r *StoryActivityRepository) Insert(activity *model.StoryActivity) (*model.StoryActivity, error) {
	if err := r.db.Create(activity).Error; err != nil {
		return nil, err
	}
	return activity, nil
}

// ListActivities for a story (paginated)
func (r *StoryActivityRepository) ListActivities(storyID uuid.UUID, lq *list.ListingQuery) ([]model.StoryActivity, int64, error) {
	query := r.db.Model(&model.StoryActivity{}).Where("story_id = ?", storyID).Order("created_at DESC")
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.StoryActivity](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
