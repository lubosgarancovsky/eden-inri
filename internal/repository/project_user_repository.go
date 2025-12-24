package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProjectUserRepository struct {
	db *gorm.DB
}

func NewProjectUserRepository(db *gorm.DB) *ProjectUserRepository {
	return &ProjectUserRepository{db}
}

func (r *ProjectUserRepository) FindAll(
	ctx context.Context,
	projectID uuid.UUID,
	lq *list.ListingQuery,
) (*[]model.ProjectUser, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(&model.ProjectUser{}).
		Preload("User").
		Where("project_id = ?", projectID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	members, total, err := helpers.List[model.ProjectUser](query, lq)
	if err != nil {
		return nil, 0, err
	}

	return &members, total, nil
}

func (r *ProjectUserRepository) FindByID(ctx context.Context, projectID, userID uuid.UUID) (*model.ProjectUser, error) {
	var projectUser model.ProjectUser
	err := r.db.
		WithContext(ctx).
		Model(&projectUser).
		Where("project_id = ? AND user_id = ?", projectID, userID).
		First(&projectUser).Error

	return &projectUser, err
}

func (r *ProjectUserRepository) Insert(
	ctx context.Context,
	projectUser *model.ProjectUser,
) (*model.ProjectUser, error) {
	err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "project_id"}, {Name: "user_id"}},
			DoNothing: true,
		}).Create(projectUser).Error

	return projectUser, err
}

func (r *ProjectUserRepository) Update(
	ctx context.Context,
	projectUser *model.ProjectUser,
) (*model.ProjectUser, error) {
	err := r.db.
		WithContext(ctx).
		Model(&projectUser).
		Clauses(clause.Returning{}).
		Where("project_id = ? AND user_id = ? AND role != ?", projectUser.ProjectID, projectUser.UserID, model.Owner).
		Updates(projectUser).Error

	return projectUser, err
}

func (r *ProjectUserRepository) Delete(
	ctx context.Context,
	memberID, projectID uuid.UUID,
) error {

	result := r.db.
		WithContext(ctx).
		Where("project_id = ? AND user_id = ? AND role != ?", projectID, memberID, model.Owner).
		Delete(&model.ProjectUser{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(
			fmt.Sprintf("User %s is not a member of project %s", memberID, projectID),
		)
	}

	return nil
}

func (r *ProjectUserRepository) GetUserRole(ctx context.Context, projectID, userID uuid.UUID) (model.ProjectRole, error) {
	var role model.ProjectRole

	err := r.db.
		WithContext(ctx).
		Model(model.ProjectUser{}).
		Select("role").
		Where("project_id = ? AND user_id = ?", projectID, userID).
		Scan(&role).Error

	return role, err
}

func (r *ProjectUserRepository) Favourite(
	ctx context.Context,
	userID, projectID uuid.UUID,
	isStarred bool,
) (*model.ProjectUser, error) {
	var projectUser model.ProjectUser
	err := r.db.
		WithContext(ctx).
		Model(&projectUser).
		Clauses(clause.Returning{}).
		Select("*").
		Where("project_id = ? AND user_id = ?", projectID, userID).
		Update("is_starred", isStarred).
		Error

	return &projectUser, err
}
