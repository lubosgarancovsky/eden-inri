package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type KanbanBoardRepository struct {
	db *gorm.DB
}

func NewKanbanBoardRepository(db *gorm.DB) *KanbanBoardRepository {
	return &KanbanBoardRepository{db: db}
}

func (r *KanbanBoardRepository) Create(ctx context.Context, board *entity.KanbanBoard) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.KanbanBoardFromDomain(board)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *KanbanBoardRepository) Update(ctx context.Context, board *entity.KanbanBoard) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ?", board.ID).Where("project_id = ?", board.ProjectID).Updates(mapper.KanbanBoardFromDomain(board))
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("board with id %s not found", board.ID))
	}
	return nil
}

func (r *KanbanBoardRepository) Delete(ctx context.Context, projectID, boardID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ?", boardID).Where("project_id = ?", projectID).Delete(&model.KanbanBoard{})
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("board with id %s not found", boardID))
	}
	return nil
}

func (r *KanbanBoardRepository) FindByID(ctx context.Context, projectID, boardID uuid.UUID) (*entity.KanbanBoard, error) {
	db := GetDB(ctx, r.db)
	var m model.KanbanBoard
	if err := db.Where("id = ?", boardID).Where("project_id = ?", projectID).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("board with id %s not found", boardID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return m.ToDomain(), nil
}

func (r *KanbanBoardRepository) List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.KanbanBoard, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.KanbanBoard{}).Where("project_id = ?", projectID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.KanbanBoard, entity.KanbanBoard](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}
