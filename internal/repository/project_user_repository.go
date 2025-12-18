package repository

import (
	"errors"
	"fmt"
	"time"

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
	return &ProjectUserRepository{db: db}
}

func (r *ProjectUserRepository) FindAll(
	userID, projectID uuid.UUID,
	lq *list.ListingQuery,
) ([]model.ProjectUser, int64, error) {

	// Step 1: ensure the caller is a member
	var count int64
	if err := r.db.
		Model(&model.ProjectUser{}).
		Where("project_id = ? AND user_id = ?", projectID, userID).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return nil, 0, api_err.ErrForbidden
	}

	// Step 2: build query for members with preloaded User
	query := r.db.
		Model(&model.ProjectUser{}).
		Preload("User").
		Where("project_id = ?", projectID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	// Step 3: apply pagination and sorting
	members, total, err := helpers.List[model.ProjectUser](query, lq)
	if err != nil {
		return nil, 0, err
	}

	return members, total, nil
}

func (r *ProjectUserRepository) Insert(
	userID, projectID uuid.UUID,
	role model.ProjectRole,
) (*model.ProjectUser, error) {

	pu := &model.ProjectUser{
		ProjectID: projectID,
		UserID:    userID,
		Role:      role,
		JoinedAt:  time.Now(),
	}

	if err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}, {Name: "user_id"}},
		DoNothing: true, // avoid duplicate membership
	}).Create(pu).Error; err != nil {
		return nil, err
	}

	return pu, nil
}

// UpdateMemberRole updates the role of a member in a project
func (r *ProjectUserRepository) Update(
	projectID, memberID uuid.UUID,
	newRole model.ProjectRole,
) (*model.ProjectUser, error) {

	// Only update role for non-owner members
	pu := &model.ProjectUser{}
	err := r.db.
		Where("project_id = ? AND user_id = ? AND role != ?", projectID, memberID, model.Owner).
		First(pu).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, api_err.ErrForbidden.WithMessage("cannot change role of owner or non-existent member")
		}
		return nil, err
	}

	pu.Role = newRole
	if err := r.db.Save(pu).Error; err != nil {
		return nil, err
	}

	return pu, nil
}

func (r *ProjectUserRepository) Delete(
	memberID, projectID uuid.UUID,
) error {

	result := r.db.
		Where("project_id = ? AND user_id = ?", projectID, memberID).
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

func (r *ProjectUserRepository) GetUserRole(projectID, userID uuid.UUID) (model.ProjectRole, error) {
	var role model.ProjectRole

	err := r.db.
		Table("inri_project_users").
		Select("role").
		Where("project_id = ? AND user_id = ?", projectID, userID).
		Scan(&role).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", api_err.ErrForbidden
		}
		return "", err
	}

	return role, nil
}

func (r *ProjectUserRepository) GetProjectUserIfMember(
	projectID, userID uuid.UUID,
) (*model.ProjectUser, error) {

	var pu model.ProjectUser

	err := r.db.
		Preload("User").
		Where("project_id = ? AND user_id = ?", projectID, userID).
		First(&pu).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, api_err.ErrForbidden.WithMessage("user is not a member of this project")
	}

	return &pu, err
}

// Update is_starred flag
func (r *ProjectUserRepository) Favourite(
	userID, projectID uuid.UUID,
	isStarred bool,
) (*model.ProjectUser, error) {
	var projectUser model.ProjectUser
	err := r.db.Model(&projectUser).
		Clauses(clause.Returning{}).
		Select("*").
		Where("project_id = ? AND user_id = ?", projectID, userID).
		Update("is_starred", isStarred).
		Error

	return &projectUser, err
}
