package model

import (
	"time"

	"github.com/google/uuid"
)

type ContactPerson struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `json:"-"`
	ClientID  uuid.UUID `json:"clientId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ContactPersonRequest struct {
	ClientID uuid.UUID `json:"clientId"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}

func (cp *ContactPerson) TableName() string {
	return "inri_contact_person"
}
