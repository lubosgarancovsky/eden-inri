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

type KanbanBoardRepository struct {
	db *gorm.DB
}

func NewKanbanBoardRepository(db *gorm.DB) *KanbanBoardRepository {
	return &KanbanBoardRepository{db: db}
}

func (r *KanbanBoardRepository) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*[]model.KanbanBoard, int64, error) {
	query := r.db.Model(model.KanbanBoard{}).
		WithContext(ctx).
		Select("*").
		Where("project_id = ?", projectID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.KanbanBoard](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *KanbanBoardRepository) Create(ctx context.Context, board *model.KanbanBoard) (*model.KanbanBoard, error) {
	if err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(board).
		Error; err != nil {
		return nil, err
	}
	return board, nil
}

func (r *KanbanBoardRepository) Update(ctx context.Context, board *model.KanbanBoard) (*model.KanbanBoard, error) {
	if err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Updates(board).
		Where("id = ? AND project_id = ?", board.ID, board.ProjectID).Error; err != nil {
		return nil, err
	}
	return board, nil
}

func (r *KanbanBoardRepository) FindByID(ctx context.Context, boardID, projectID uuid.UUID) (*model.KanbanBoard, error) {
	var board model.KanbanBoard
	err := r.db.
		WithContext(ctx).
		Where("id = ? AND project_id = ?", boardID, projectID).
		First(&board).
		Error
	return &board, err
}

func (r *KanbanBoardRepository) Delete(ctx context.Context, boardID, projectID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("id = ? AND project_id = ?", boardID, projectID).
		Delete(&model.KanbanBoard{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}

func (r *KanbanBoardRepository) GetProjectIDByBoardID(ctx context.Context, boardID uuid.UUID) (uuid.UUID, error) {
	var projectID uuid.UUID
	err := r.db.
		WithContext(ctx).
		Where("id = ?", boardID).
		Select("project_id").
		First(&projectID).
		Error

	return projectID, err
}
