package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (r *ProjectRepository) List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Project, int64, error) {
	db := GetDB(ctx, r.db)

	query := db.Table("inri_projects p").
		Select("p.*, pu.role, pu.is_starred").
		Joins("JOIN inri_project_users pu ON pu.project_id = p.id").
		Where("pu.user_id = ?", userID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.Project, entity.Project](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *ProjectRepository) FindByID(ctx context.Context, userID, projectID uuid.UUID) (*entity.Project, error) {
	db := GetDB(ctx, r.db)

	var prj model.Project
	err := db.Table("inri_projects p").
		Select("p.*, pu.role, pu.is_starred").
		Joins("JOIN inri_project_users pu ON pu.project_id = p.id").
		Where("p.id = ? AND pu.user_id = ?", projectID, userID).
		First(&prj).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("project with id %s not found", projectID))
		}

		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	return prj.ToDomain(), nil
}

func (r *ProjectRepository) Create(ctx context.Context, project *entity.Project) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.ProjectFromDomain(project)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *ProjectRepository) Update(ctx context.Context, project *entity.Project) error {
	db := GetDB(ctx, r.db)
	md := mapper.ProjectFromDomain(project)

	result := db.
		Where("id", project.ID).
		Updates(md)

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("project with id %s not found", project.ID))
	}

	return nil
}

func (r *ProjectRepository) Delete(ctx context.Context, projectID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("id", projectID).
		Delete(&model.Project{})

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("project with id %s not found", projectID))
	}

	return nil
}
