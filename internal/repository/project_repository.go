package repository

import (
	"context"
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

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) WithTx(ctx context.Context, fn func(txRepo *ProjectRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := &ProjectRepository{db: tx}
		return fn(txRepo)
	})
}

func (r *ProjectRepository) FindAll(
	ctx context.Context,
	userID uuid.UUID,
	lq *list.ListingQuery,
) (*[]model.Project, int64, error) {

	query := r.db.
		WithContext(ctx).
		Table("inri_projects p").
		Select("p.*, pu.role, pu.is_starred").
		Joins(`
			JOIN inri_project_users pu
			  ON pu.project_id = p.id
		`).
		Where("pu.user_id = ?", userID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.Project](query, lq)
	if err != nil {
		return nil, 0, err
	}

	return &items, total, nil
}

func (r *ProjectRepository) FindByID(
	ctx context.Context,
	userID,
	projectID uuid.UUID,

) (*model.Project, error) {
	var result model.Project

	err := r.db.
		WithContext(ctx).
		Table("inri_projects p").
		Select("p.*, pu.role, pu.is_starred").
		Joins(`
			JOIN inri_project_users pu
			  ON pu.project_id = p.id
		`).
		Where("p.id = ? AND pu.user_id = ?", projectID, userID).
		First(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ProjectRepository) InsertProject(
	ctx context.Context,
	prj *model.Project,
) (*model.Project, error) {
	if err := r.db.
		WithContext(ctx).
		Model(&prj).
		Clauses(clause.Returning{}).
		Select("*").
		Create(&prj).
		Error; err != nil {
		return nil, err
	}

	return prj, nil
}

func (r *ProjectRepository) InsertProjectUser(
	ctx context.Context,
	pu *model.ProjectUser,
) (*model.ProjectUser, error) {

	if err := r.db.Model(&pu).
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Select("*").
		Create(&pu).
		Error; err != nil {
		return nil, err
	}

	return pu, nil
}

func (r *ProjectRepository) Update(ctx context.Context, prj *model.Project) (*model.Project, error) {
	result := r.db.
		WithContext(ctx).
		Model(&prj).
		Clauses(clause.Returning{}).
		Select("*").
		Where("id = ?", prj.ID).
		Updates(map[string]interface{}{
			"name":             prj.Name,
			"description":      prj.Description,
			"tags":             prj.Tags,
			"status":           prj.Status,
			"updated_at":       prj.UpdatedAt,
			"last_activity_at": prj.LastActivityAt,
		})
	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Project with id %s does not exist", prj.ID))
	}
	return prj, nil
}

func (r *ProjectRepository) Delete(ctx context.Context, projectID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("id = ?", projectID).
		Delete(model.Project{})

	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Project with id %s does not exist", projectID))
	}
	return nil
}

func (r *ProjectRepository) ListProjectMembers(
	ctx context.Context,
	userID, projectID uuid.UUID,
	lq *list.ListingQuery,
) ([]model.ProjectUser, int64, error) {

	// Ensure the caller is a member
	var count int64
	if err := r.db.
		WithContext(ctx).
		Model(&model.ProjectUser{}).
		Where("project_id = ? AND user_id = ?", projectID, userID).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return nil, 0, api_err.ErrForbidden
	}

	// Build query for members with preloaded User
	query := r.db.
		WithContext(ctx).
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

func (r *ProjectRepository) AddMember(
	ctx context.Context,
	userID, projectID uuid.UUID,
	role model.ProjectRole,
) (*model.ProjectUser, error) {

	pu := &model.ProjectUser{
		ProjectID: projectID,
		UserID:    userID,
		Role:      role,
		JoinedAt:  time.Now(),
	}

	if err := r.db.
		WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "project_id"}, {Name: "user_id"}},
			DoNothing: true, // avoid duplicate membership
		}).Create(pu).Error; err != nil {
		return nil, err
	}

	return pu, nil
}

func (r *ProjectRepository) RemoveMember(
	ctx context.Context,
	memberID, projectID uuid.UUID,
) error {

	result := r.db.
		WithContext(ctx).
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

func (r *ProjectRepository) GetUserRole(ctx context.Context, projectID, userID uuid.UUID) (model.ProjectRole, error) {
	var role model.ProjectRole

	err := r.db.
		WithContext(ctx).
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

// UpdateMemberRole updates the role of a member in a project
func (r *ProjectRepository) UpdateMemberRole(
	ctx context.Context,
	projectID, memberID uuid.UUID,
	newRole model.ProjectRole,
) (*model.ProjectUser, error) {

	// Only update role for non-owner members
	pu := &model.ProjectUser{}
	err := r.db.
		WithContext(ctx).
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
