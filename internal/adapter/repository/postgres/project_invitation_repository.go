package postgres

import (
	"context"
	"errors"

	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type ProjectInvitationRepository struct {
	db *gorm.DB
}

func NewProjectInvitationRepository(db *gorm.DB) *ProjectInvitationRepository {
	return &ProjectInvitationRepository{db: db}
}

func (r *ProjectInvitationRepository) FindByToken(ctx context.Context, token string) (*entity.ProjectInvitation, error) {
	db := GetDB(ctx, r.db)

	var invitation model.ProjectInvitation
	err := db.Where("token = ?", token).First(&invitation).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, go_kit.ErrNotFound.WithMessage("invitation not found")
		}

		return nil, err
	}

	return invitation.ToDomain(), nil
}

func (r *ProjectInvitationRepository) Save(ctx context.Context, invitation *entity.ProjectInvitation) error {
	db := GetDB(ctx, r.db)

	result := db.Save(mapper.ProjectInvitationFromDomain(invitation))

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	return nil
}
