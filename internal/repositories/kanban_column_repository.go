package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
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
func (r *KanbanColumnRepository) FindAll(ctx context.Context, boardID uuid.UUID) ([]models.KanbanColumn, error) {
	var cols []models.KanbanColumn
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
func (r *KanbanColumnRepository) FindByID(ctx context.Context, boardID, columnID uuid.UUID) (*models.KanbanColumn, error) {
	var col models.KanbanColumn
	err := r.db.
		WithContext(ctx).
		Where("board_id = ? AND id = ?", boardID, columnID).
		First(&col).
		Error

	return &col, err
}

// Insert a new column
func (r *KanbanColumnRepository) Insert(ctx context.Context, col *models.KanbanColumn) (*models.KanbanColumn, error) {
	col.ID = uuid.New()
	col.CreatedAt = time.Now()

	err := r.db.
		WithContext(ctx).
		Create(col).
		Error

	return col, err
}

// Update column
func (r *KanbanColumnRepository) Update(ctx context.Context, col *models.KanbanColumn) (*models.KanbanColumn, error) {
	return col, r.db.WithContext(ctx).Save(col).Error
}

// Delete column
func (r *KanbanColumnRepository) Delete(ctx context.Context, boardID, columnID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("board_id = ? AND id = ?", boardID, columnID).
		Delete(models.KanbanColumn{})

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
	var col models.KanbanColumn
	err := r.db.
		WithContext(ctx).
		Where("id = ?", columnID).
		First(&col).
		Error

	return &col.BoardID, err
}
