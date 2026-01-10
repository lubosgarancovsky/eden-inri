package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
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
) (*[]models.ProjectUser, int64, error) {
	query := r.db.WithContext(ctx).
		Model(&models.ProjectUser{}).
		Preload("User").
		Joins(`JOIN iam_users u ON u.id = inri_project_users.user_id`).
		Where("project_id = ?", projectID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	members, total, err := helpers.List[models.ProjectUser](query, lq)
	if err != nil {
		return nil, 0, err
	}

	return &members, total, nil
}

func (r *ProjectUserRepository) FindByID(ctx context.Context, userID, projectID uuid.UUID) (*models.ProjectUser, error) {
	var projectUser models.ProjectUser
	err := r.db.
		WithContext(ctx).
		Model(&projectUser).
		Preload("User").
		Where("project_id = ? AND user_id = ?", projectID, userID).
		First(&projectUser).Error

	return &projectUser, err
}

func (r *ProjectUserRepository) Insert(
	ctx context.Context,
	projectUser *models.ProjectUser,
) (*models.ProjectUser, error) {
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
	projectUser *models.ProjectUser,
) (*models.ProjectUser, error) {
	err := r.db.
		WithContext(ctx).
		Model(&projectUser).
		Clauses(clause.Returning{}).
		Where("project_id = ? AND user_id = ? AND role != ?", projectUser.ProjectID, projectUser.UserID, models.Owner).
		Updates(projectUser).Error

	return projectUser, err
}

func (r *ProjectUserRepository) Delete(
	ctx context.Context,
	memberID, projectID uuid.UUID,
) error {

	result := r.db.
		WithContext(ctx).
		Where("project_id = ? AND user_id = ? AND role != ?", projectID, memberID, models.Owner).
		Delete(&models.ProjectUser{})

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

func (r *ProjectUserRepository) GetUserRole(ctx context.Context, projectID, userID uuid.UUID) (models.ProjectRole, error) {
	var role models.ProjectRole

	err := r.db.
		WithContext(ctx).
		Model(models.ProjectUser{}).
		Select("role").
		Where("project_id = ? AND user_id = ?", projectID, userID).
		Scan(&role).Error

	return role, err
}

func (r *ProjectUserRepository) Favourite(
	ctx context.Context,
	userID, projectID uuid.UUID,
	isStarred bool,
) (*models.ProjectUser, error) {
	var projectUser models.ProjectUser
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

func (r *ProjectUserRepository) GetOwner(ctx context.Context, projectID uuid.UUID) (*models.ProjectUser, error) {
	var projectUser models.ProjectUser
	err := r.db.WithContext(ctx).
		Where("project_id = ? AND role = ?", projectID, models.Owner).
		First(&projectUser).
		Error

	return &projectUser, err
}
