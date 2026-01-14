package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type ProjectInvitation struct {
	ID         uuid.UUID          `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	ProjectID  uuid.UUID          `gorm:"type:uuid;not null"`
	UserID     uuid.UUID          `gorm:"type:uuid;not null"`
	Role       entity.ProjectRole `gorm:"type:text;not null"`
	Token      string             `gorm:"type:text;not null"`
	InvitedBy  uuid.UUID          `gorm:"type:uuid;not null"`
	CreatedAt  time.Time          `gorm:"type:timestamptz;autoCreateTime;not null"`
	ExpiresAt  time.Time          `gorm:"type:timestamptz;autoCreateTime;not null"`
	AcceptedAt *time.Time         `gorm:"type:timestamptz;autoCreateTime"`
}

func (i *ProjectInvitation) ToDomain() *entity.ProjectInvitation {
	return &entity.ProjectInvitation{
		ID:         i.ID,
		ProjectID:  i.ProjectID,
		UserID:     i.UserID,
		Role:       i.Role,
		Token:      i.Token,
		InvitedBy:  i.InvitedBy,
		CreatedAt:  i.CreatedAt,
		ExpiresAt:  i.ExpiresAt,
		AcceptedAt: i.AcceptedAt,
	}
}

func (c *ProjectInvitation) TableName() string {
	return "inri_project_invitations"
}
