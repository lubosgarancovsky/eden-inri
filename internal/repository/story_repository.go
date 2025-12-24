package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/lubosgarancovsky/eden-inri/internal/model"
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

func (r *StoryRepository) FindAll(ctx context.Context, boardID, columnID uuid.UUID, lq *list.ListingQuery) (*[]model.Story, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(model.Story{}).
		Where("board_id = ? AND column_id = ?", boardID, columnID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.Story](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil

}

func (r *StoryRepository) FindByID(ctx context.Context, boardID, storyID uuid.UUID) (*model.Story, error) {
	var story model.Story
	err := r.db.
		WithContext(ctx).
		Where("board_id = ? AND id = ?", boardID, storyID).
		First(&story).
		Error

	return &story, err
}

func (r *StoryRepository) Insert(ctx context.Context, story *model.Story) (*model.Story, error) {
	err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(story).
		Error

	return story, err
}

func (r *StoryRepository) Update(ctx context.Context, story *model.Story) (*model.Story, error) {
	err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("board_id = ? AND column_id = ? AND id = ?", story.BoardID, story.ColumnID, story.ID).
		Updates(story).
		Error

	return story, err
}

func (r *StoryRepository) Delete(ctx context.Context, boardID, storyID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("board_id = ? AND id = ?", boardID, storyID).
		Delete(model.Story{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}

	return nil
}
