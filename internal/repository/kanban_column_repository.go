package repository

import (
	"errors"
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
func (r *KanbanColumnRepository) FindAll(boardID uuid.UUID) ([]model.KanbanColumn, error) {
	var cols []model.KanbanColumn
	if err := r.db.Where("board_id = ?", boardID).Order("position ASC").Find(&cols).Error; err != nil {
		return nil, err
	}
	return cols, nil
}

// FindByID a single column
func (r *KanbanColumnRepository) FindByID(columnID uuid.UUID) (*model.KanbanColumn, error) {
	var col model.KanbanColumn
	err := r.db.Where("id = ?", columnID).First(&col).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, api_err.ErrNotFound
	}
	return &col, err
}

// Insert a new column
func (r *KanbanColumnRepository) Insert(col *model.KanbanColumn) (*model.KanbanColumn, error) {
	col.ID = uuid.New()
	col.CreatedAt = time.Now()

	if err := r.db.Create(col).Error; err != nil {
		return nil, err
	}
	return col, nil
}

// Update column
func (r *KanbanColumnRepository) Update(col *model.KanbanColumn) (*model.KanbanColumn, error) {
	if err := r.db.Save(col).Error; err != nil {
		return nil, err
	}
	return col, nil
}

// Delete column
func (r *KanbanColumnRepository) Delete(columnID uuid.UUID) error {
	result := r.db.Where("id = ?", columnID).Delete(&model.KanbanColumn{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}
