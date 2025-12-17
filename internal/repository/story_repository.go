package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"

	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

type StoryRepository struct {
	db *gorm.DB
}

func NewStoryRepository(db *gorm.DB) *StoryRepository {
	return &StoryRepository{db: db}
}

func (r *StoryRepository) DB() *gorm.DB {
	return r.db
}

// FindAll stories for a column
func (r *StoryRepository) FindAll(columnID uuid.UUID, lq *list.ListingQuery) ([]model.Story, int64, error) {
	query := r.db.Model(&model.Story{}).Where("column_id = ?", columnID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.Story](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil

}

// FindByID a single story
func (r *StoryRepository) FindByID(storyID uuid.UUID) (*model.Story, error) {
	var story model.Story
	err := r.db.Preload("Labels").Preload("Activity").
		Where("id = ?", storyID).First(&story).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, api_err.ErrNotFound
	}
	return &story, err
}

// Insert a new story
func (r *StoryRepository) Insert(story *model.Story) (*model.Story, error) {
	story.ID = uuid.New()
	story.CreatedAt = time.Now()
	if err := r.db.Create(story).Error; err != nil {
		return nil, err
	}
	return story, nil
}

// Update story
func (r *StoryRepository) Update(story *model.Story) (*model.Story, error) {
	if err := r.db.Save(story).Error; err != nil {
		return nil, err
	}
	return story, nil
}

// Delete story
func (r *StoryRepository) Delete(storyID uuid.UUID) error {
	result := r.db.Where("id = ?", storyID).Delete(&model.Story{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}
