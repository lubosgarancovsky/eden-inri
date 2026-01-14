package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type ProjectUserRepository struct {
	db *gorm.DB
}

func NewProjectUserRepository(db *gorm.DB) *ProjectUserRepository {
	return &ProjectUserRepository{
		db: db,
	}
}

func (r *ProjectUserRepository) List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ProjectUser, int64, error) {
	db := GetDB(ctx, r.db)

	query := db.Model(&model.ProjectUser{}).Preload("User").Where("project_id = ?", projectID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.ProjectUser, entity.ProjectUser](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *ProjectUserRepository) Delete(ctx context.Context, projectID, memberID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	result := db.Where("project_id = ? AND user_id = ?", projectID, memberID).Delete(&model.ProjectUser{})
	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("member with id %s not found in project %s", memberID, projectID))
	}

	return nil
}

func (r *ProjectUserRepository) FindByID(ctx context.Context, projectID, memberID uuid.UUID) (*entity.ProjectUser, error) {
	db := GetDB(ctx, r.db)
	var pu model.ProjectUser
	if err := db.
		Where("project_id = ? AND user_id = ?", projectID, memberID).
		Preload("User").
		First(&pu).Error; err != nil {
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	return pu.ToDomain(), nil
}

func (r *ProjectUserRepository) SetRole(ctx context.Context, projectID, memberID uuid.UUID, role entity.ProjectRole) error {
	db := GetDB(ctx, r.db)
	result := db.Model(&model.ProjectUser{}).Where("project_id = ? AND user_id = ?", projectID, memberID).Update("role", string(role))
	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("member with id %s not found in project %s", memberID, projectID))
	}

	return nil
}

func (r *ProjectUserRepository) Insert(ctx context.Context, pu *entity.ProjectUser) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.ProjectUserFromDomain(pu)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *ProjectUserRepository) Favourite(ctx context.Context, userID, projectID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	var pu model.ProjectUser

	if err := db.Where("project_id = ? AND user_id = ?", projectID, userID).First(&pu).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	pu.IsStarred = !pu.IsStarred

	if err := db.Save(&pu).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	return nil
}

func (r *ProjectUserRepository) IsMember(ctx context.Context, userID, projectID uuid.UUID) (bool, error) {
	db := GetDB(ctx, r.db)
	var pu model.ProjectUser

	if err := db.Where("project_id = ? AND user_id = ?", projectID, userID).First(&pu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	return true, nil
}

func (r *ProjectUserRepository) HasRole(ctx context.Context, userID, projectID uuid.UUID, roles []entity.ProjectRole) (bool, error) {
	db := GetDB(ctx, r.db)
	var pu model.ProjectUser

	if err := db.Where("project_id = ? AND user_id = ?", projectID, userID).First(&pu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, app_err.ErrNotAMember
		}

		return false, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	for _, role := range roles {
		if pu.Role == string(role) {
			return true, nil
		}
	}

	return false, nil
}

func (r *ProjectUserRepository) GetOwner(ctx context.Context, projectID uuid.UUID) (uuid.UUID, error) {
	db := GetDB(ctx, r.db)
	var pu model.ProjectUser

	if err := db.Where("project_id = ? AND role = ?", projectID, "owner").First(&pu).Error; err != nil {
		return uuid.Nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	return pu.UserID, nil
}
