package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"gorm.io/gorm"
)

type KanbanColumnRepository struct {
	db *gorm.DB
}

func NewKanbanColumnRepository(db *gorm.DB) *KanbanColumnRepository {
	return &KanbanColumnRepository{db: db}
}

// FindAll columns of a board
func (r *KanbanColumnRepository) FindAll(ctx context.Context, boardID uuid.UUID) ([]model.KanbanColumn, error) {
	var cols []model.KanbanColumn
	if err := r.db.
		WithContext(ctx).
		Where("board_id = ?", boardID).
		Order("position ASC").Find(&cols).
		Error; err != nil {
		return nil, err
	}
	return cols, nil
}

// FindByID a single column
func (r *KanbanColumnRepository) FindByID(ctx context.Context, boardID, columnID uuid.UUID) (*model.KanbanColumn, error) {
	var col model.KanbanColumn
	err := r.db.
		WithContext(ctx).
		Where("board_id = ? AND id = ?", boardID, columnID).
		First(&col).
		Error

	return &col, err
}

// Insert a new column
func (r *KanbanColumnRepository) Insert(ctx context.Context, col *model.KanbanColumn) (*model.KanbanColumn, error) {
	col.ID = uuid.New()
	col.CreatedAt = time.Now()

	err := r.db.
		WithContext(ctx).
		Create(col).
		Error

	return col, err
}

// Update column
func (r *KanbanColumnRepository) Update(ctx context.Context, col *model.KanbanColumn) (*model.KanbanColumn, error) {
	return col, r.db.WithContext(ctx).Save(col).Error
}

// Delete column
func (r *KanbanColumnRepository) Delete(ctx context.Context, boardID, columnID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("board_id = ? AND id = ?", boardID, columnID).
		Delete(model.KanbanColumn{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}

// GetBoardID returns a board ID of a column
func (r *KanbanColumnRepository) GetBoardID(ctx context.Context, columnID uuid.UUID) (*uuid.UUID, error) {
	var boardID uuid.UUID
	err := r.db.
		WithContext(ctx).
		Model(&model.KanbanColumn{}).
		Where("id = ?", columnID).
		Select("board_id").
		First(&boardID).
		Error

	return &boardID, err
}
