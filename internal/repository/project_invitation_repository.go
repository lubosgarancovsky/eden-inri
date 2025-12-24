package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProjectInvitationRepository struct {
	db *gorm.DB
}

func NewProjectInvitationRepository(db *gorm.DB) *ProjectInvitationRepository {
	return &ProjectInvitationRepository{db: db}
}

func (r *ProjectInvitationRepository) FindByToken(ctx context.Context, userID uuid.UUID, token string) (*model.ProjectInvitation, error) {
	var invitation model.ProjectInvitation
	if err := r.db.
		WithContext(ctx).
		Model(model.ProjectInvitation{}).
		Select("*").
		Where("token = ? AND user_id", token, userID).
		First(&invitation).Error; err != nil {
		return nil, err
	}
	return &invitation, nil
}

func (r *ProjectInvitationRepository) Insert(ctx context.Context, invitation *model.ProjectInvitation) (*model.ProjectInvitation, error) {
	if err := r.db.
		WithContext(ctx).
		Model(&invitation).
		Clauses(clause.Returning{}).
		Select("*").
		Create(invitation).
		Error; err != nil {
		return nil, err
	}

	return invitation, nil
}

func (r *ProjectInvitationRepository) Update(ctx context.Context, invitation *model.ProjectInvitation) (*model.ProjectInvitation, error) {
	result := r.db.
		WithContext(ctx).
		Model(&invitation).
		Clauses(clause.Returning{}).
		Select("*").
		Where("user_id = ? AND id = ?", invitation.UserID, invitation.ID).
		Updates(invitation)

	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage("Project invitation was not found")
	}
	return invitation, nil
}

func (r *ProjectInvitationRepository) Delete(ctx context.Context, userID, invitationID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("user_id = ? AND id = ?", userID, invitationID).
		Delete(model.ProjectInvitation{})

	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage("Project invitation was not found")
	}

	return nil
}

func (r *ProjectInvitationRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.Model(&user).
		WithContext(ctx).
		Select("id, username, email, first_name, last_name").
		Where("email = ?", email).
		First(&user).
		Error

	return &user, err
}
