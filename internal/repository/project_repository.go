package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) FindAll(userID uuid.UUID, lq *list.ListingQuery) ([]model.Project, int64, error) {
	query := r.db.Model(&model.Project{}).Where("user_id = ?", userID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.Project](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *ProjectRepository) FindByID(id uuid.UUID) (*model.Project, error) {
	var result model.Project
	if err := r.db.Model(&model.Project{}).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ProjectRepository) Insert(prj *model.Project) (*model.Project, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(prj).Error; err != nil {
		return nil, err
	}
	return prj, nil
}

func (r *ProjectRepository) Update(prj *model.Project) (*model.Project, error) {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", prj.ID).Updates(&prj)
	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Project with id %s does not exist", prj.ID))
	}
	return prj, nil
}

func (r *ProjectRepository) Delete(id uuid.UUID) error {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(model.Project{})
	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Project with id %s does not exist", id))
	}
	return nil
}
