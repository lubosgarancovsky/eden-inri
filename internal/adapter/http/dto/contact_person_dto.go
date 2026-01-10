package dto

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

type CreateContactPersonReq struct {
	ClientID uuid.UUID `uri:"clientId"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}

type UpdateContactPersonReq struct {
	ClientID        uuid.UUID `uri:"clientId"`
	ContactPersonID uuid.UUID `uri:"contactPersonId"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
}
