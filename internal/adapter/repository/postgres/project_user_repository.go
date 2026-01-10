package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	go_kit "github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

var _ ports.ProjectUserPersistPort = (*ProjectUserRepository)(nil)

type ProjectUserRepository struct {
	db *gorm.DB
}

func NewProjectUserRepository(db *gorm.DB) *ProjectUserRepository {
	return &ProjectUserRepository{db: db}
}

type projectUserRow struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	Role      string
	IsStarred bool
}

func (r *ProjectUserRepository) AddOwner(ctx context.Context, userID, projectID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	row := &projectUserRow{ProjectID: projectID, UserID: userID, Role: "owner", IsStarred: false}
	if err := db.Table("inri_project_users").Create(row).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}
