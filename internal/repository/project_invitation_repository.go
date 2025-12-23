package repository

import (
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

func (r *ProjectInvitationRepository) FindByToken(userID uuid.UUID, token string) (*model.ProjectInvitation, error) {
	var invitation model.ProjectInvitation
	if err := r.db.
		Model(model.ProjectInvitation{}).
		Select("*").
		Where("token = ? AND user_id", token, userID).
		First(&invitation).Error; err != nil {
		return nil, err
	}
	return &invitation, nil
}

func (r *ProjectInvitationRepository) Insert(invitation *model.ProjectInvitation) (*model.ProjectInvitation, error) {
	if err := r.db.
		Model(&invitation).
		Clauses(clause.Returning{}).
		Select("*").
		Create(invitation).
		Error; err != nil {
		return nil, err
	}

	return invitation, nil
}

func (r *ProjectInvitationRepository) Update(invitation *model.ProjectInvitation) (*model.ProjectInvitation, error) {
	result := r.db.
		Model(&invitation).
		Clauses(clause.Returning{}).
		Select("*").
		Updates(invitation)

	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage("Project invitation was not found")
	}
	return invitation, nil
}

func (r *ProjectInvitationRepository) Delete(invitation *model.ProjectInvitation) error {
	result := r.db.
		Model(&invitation).
		Delete(invitation)
	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage("Project invitation was not found")
	}

	return nil
}

func (r *ProjectInvitationRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Model(&user).
		Select("id, username, email, first_name, last_name").
		Where("email = ?", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}
