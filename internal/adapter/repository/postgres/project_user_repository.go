package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
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
			return false, nil
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
