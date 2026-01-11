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

type ProjectLabelRepository struct {
	db *gorm.DB
}

func NewProjectLabelRepository(db *gorm.DB) *ProjectLabelRepository {
	return &ProjectLabelRepository{db: db}
}

func (r *ProjectLabelRepository) Create(ctx context.Context, label *entity.ProjectLabel) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.ProjectLabelFromDomain(label)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *ProjectLabelRepository) Update(ctx context.Context, label *entity.ProjectLabel) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ?", label.ID).Where("project_id = ?", label.ProjectID).Updates(mapper.ProjectLabelFromDomain(label))
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("label with id %s not found", label.ID))
	}
	return nil
}

func (r *ProjectLabelRepository) Delete(ctx context.Context, projectID, labelID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ?", labelID).Where("project_id = ?", projectID).Delete(&model.ProjectLabel{})
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("label with id %s not found", labelID))
	}
	return nil
}

func (r *ProjectLabelRepository) FindByID(ctx context.Context, projectID, labelID uuid.UUID) (*entity.ProjectLabel, error) {
	db := GetDB(ctx, r.db)
	var m model.ProjectLabel
	if err := db.Where("id = ?", labelID).Where("project_id = ?", projectID).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("label with id %s not found", labelID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return m.ToDomain(), nil
}

func (r *ProjectLabelRepository) List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ProjectLabel, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.ProjectLabel{}).Where("project_id = ?", projectID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.ProjectLabel, entity.ProjectLabel](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}
