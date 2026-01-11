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

type KanbanColumnRepository struct {
	db *gorm.DB
}

func NewKanbanColumnRepository(db *gorm.DB) *KanbanColumnRepository {
	return &KanbanColumnRepository{db: db}
}

func (r *KanbanColumnRepository) Create(ctx context.Context, col *entity.KanbanColumn) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.KanbanColumnFromDomain(col)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *KanbanColumnRepository) Update(ctx context.Context, col *entity.KanbanColumn) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ?", col.ID).Where("board_id = ?", col.BoardID).Updates(mapper.KanbanColumnFromDomain(col))
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("column with id %s not found", col.ID))
	}
	return nil
}

func (r *KanbanColumnRepository) Delete(ctx context.Context, boardID, columnID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ?", columnID).Where("board_id = ?", boardID).Delete(&model.KanbanColumn{})
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("column with id %s not found", columnID))
	}
	return nil
}

func (r *KanbanColumnRepository) FindByID(ctx context.Context, boardID, columnID uuid.UUID) (*entity.KanbanColumn, error) {
	db := GetDB(ctx, r.db)
	var m model.KanbanColumn
	if err := db.Where("id = ?", columnID).Where("board_id = ?", boardID).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("column with id %s not found", columnID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return m.ToDomain(), nil
}

func (r *KanbanColumnRepository) List(ctx context.Context, boardID uuid.UUID) ([]entity.KanbanColumn, error) {
	db := GetDB(ctx, r.db)
	var items []model.KanbanColumn
	if err := db.Where("board_id = ?", boardID).Order("position asc").Find(&items).Error; err != nil {
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	result := make([]entity.KanbanColumn, len(items))
	for i, item := range items {
		result[i] = *item.ToDomain()
	}
	return result, nil
}
