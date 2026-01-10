package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

type StoryRepository struct {
	db *gorm.DB
}

func NewStoryRepository(db *gorm.DB) *StoryRepository {
	return &StoryRepository{db: db}
}

func (r *StoryRepository) WithTx(ctx context.Context, fn func(txRepo *StoryRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := &StoryRepository{db: tx}
		return fn(txRepo)
	})
}

func (r *StoryRepository) DB() *gorm.DB {
	return r.db
}

func (r *StoryRepository) FindAllAssigned(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]models.StoryListItem, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(models.Story{}).
		Preload("Assignee").
		Select("*").
		Where("assignee_id = ?", userID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[models.StoryListItem](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *StoryRepository) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*[]models.StoryListItem, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(models.Story{}).
		Preload("Assignee").
		Select("*").
		Where("project_id = ?", projectID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[models.StoryListItem](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil

}

func (r *StoryRepository) FindByID(ctx context.Context, projectID, storyID uuid.UUID) (*models.Story, error) {
	var story models.Story
	err := r.db.
		WithContext(ctx).
		Preload("Assignee").
		Preload("Creator").
		Select("*").
		Where("project_id = ? AND id = ?", projectID, storyID).
		First(&story).
		Error

	return &story, err
}

func (r *StoryRepository) FindBySlug(ctx context.Context, projectID uuid.UUID, slug string) (*models.Story, error) {
	var story models.Story
	err := r.db.
		WithContext(ctx).
		Preload("Assignee").
		Preload("Creator").
		Select("*").
		Where("project_id = ? AND slug = ?", projectID, slug).
		First(&story).
		Error

	return &story, err
}

func (r *StoryRepository) Insert(ctx context.Context, story *models.Story) (*models.Story, error) {
	err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Select("*").
		Create(story).
		Error

	return story, err
}

func (r *StoryRepository) Update(ctx context.Context, story *models.Story) (*models.Story, error) {
	err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("id = ?", story.ID).
		Updates(story).
		Error

	return story, err
}

func (r *StoryRepository) Delete(ctx context.Context, projectID, storyID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("project_id = ? AND id = ?", projectID, storyID).
		Delete(models.Story{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}

	return nil
}

func (r *StoryRepository) GetAssignee(ctx context.Context, projectID, storyID uuid.UUID) (*models.User, error) {
	var story models.Story
	err := r.db.
		WithContext(ctx).
		Preload("Assignee").
		Preload("Creator").
		Select("*").
		Where("project_id = ? AND id = ?", projectID, storyID).
		First(&story).
		Error

	return story.Assignee, err
}

func (r *StoryRepository) ChangeAssignee(ctx context.Context, projectID, storyID uuid.UUID, assigneeID *uuid.UUID) error {
	return r.db.
		WithContext(ctx).
		Model(models.Story{}).
		Select("*").
		Where("project_id = ? AND id = ?", projectID, storyID).
		Update("assignee_id", assigneeID).
		Error
}
