package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

type KanbanBoardRepository struct {
	db *gorm.DB
}

func NewKanbanBoardRepository(db *gorm.DB) *KanbanBoardRepository {
	return &KanbanBoardRepository{db: db}
}

// List boards of a project
func (r *KanbanBoardRepository) FindAll(projectID uuid.UUID) ([]model.KanbanBoard, error) {
	var boards []model.KanbanBoard
	if err := r.db.Where("project_id = ?", projectID).Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}

// Create a new Kanban board
func (r *KanbanBoardRepository) Create(board *model.KanbanBoard) (*model.KanbanBoard, error) {
	board.ID = uuid.New()
	board.CreatedAt = time.Now()

	if err := r.db.Create(board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

// Find a board by ID, ensuring it belongs to the project
func (r *KanbanBoardRepository) FindByID(boardID, projectID uuid.UUID) (*model.KanbanBoard, error) {
	var board model.KanbanBoard
	err := r.db.Where("id = ? AND project_id = ?", boardID, projectID).First(&board).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, api_err.ErrNotFound
	}
	return &board, err
}

// Delete a board
func (r *KanbanBoardRepository) Delete(boardID, projectID uuid.UUID) error {
	result := r.db.Where("id = ? AND project_id = ?", boardID, projectID).Delete(&model.KanbanBoard{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}
