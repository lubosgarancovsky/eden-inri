package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"
)

type ContactPerson struct {
	ID        uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null"`
	ClientID  uuid.UUID      `gorm:"type:uuid;not null"`
	Name      string         `gorm:"type:string;not null"`
	Email     *string        `gorm:"type:string"`
	Phone     *string        `gorm:"type:string"`
	CreatedAt time.Time      `gorm:"type:timestamptz;autoCreateTime;not null"`
	UpdatedAt time.Time      `gorm:"type:timestamptz;autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz;index"`
}

func (ContactPerson) TableName() string {
	return "inri_contact_person"
}

func (c ContactPerson) ToDomain() *entity.ContactPerson {
	return &entity.ContactPerson{
		ID:        c.ID,
		UserID:    c.UserID,
		ClientID:  c.ClientID,
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
