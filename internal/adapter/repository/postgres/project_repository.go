package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

var _ ports.QueryProjectPort = (*ProjectRepository)(nil)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Project, int64, error) {
	db := GetDB(ctx, r.db)

	query := db.Model(&model.Project{}).
		Joins("JOIN inri_project_users pu ON pu.project_id = inri_projects.id AND pu.user_id = ?", userID).
		Select("inri_projects.*, pu.is_starred as is_starred, pu.role as role")

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := List[entity.Project](query, lq)
	if err != nil {
		return nil, 0, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return &items, total, nil
}

func (r *ProjectRepository) FindByID(ctx context.Context, userID, projectID uuid.UUID) (*entity.Project, error) {
	db := GetDB(ctx, r.db)

	var m model.Project
	if err := db.Model(&model.Project{}).
		Joins("JOIN inri_project_users pu ON pu.project_id = inri_projects.id AND pu.user_id = ?", userID).
		Select("inri_projects.*, pu.is_starred as is_starred, pu.role as role").
		Where("inri_projects.id = ?", projectID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("project with id %s not found", projectID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return m.ToDomain(), nil
}

func (r *ProjectRepository) Create(ctx context.Context, project *entity.Project) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(model.ProjectFromDomain(project)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *ProjectRepository) Update(ctx context.Context, userID uuid.UUID, project *entity.Project) error {
	db := GetDB(ctx, r.db)

	result := db.Model(&model.Project{}).
		Joins("JOIN inri_project_users pu ON pu.project_id = inri_projects.id AND pu.user_id = ?", userID).
		Where("inri_projects.id = ?", project.ID).
		Updates(model.ProjectFromDomain(project))

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("project with id %s not found", project.ID))
	}
	return nil
}

func (r *ProjectRepository) Delete(ctx context.Context, userID, projectID uuid.UUID) error {
	db := GetDB(ctx, r.db)

	result := db.Table("inri_projects p").
		Joins("JOIN inri_project_users pu ON pu.project_id = p.id AND pu.user_id = ? AND pu.role = ?", userID, "owner").
		Where("p.id = ?", projectID).
		Delete(&model.Project{})

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("project with id %s not found", projectID))
	}
	return nil
}
