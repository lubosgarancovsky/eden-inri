package entity

import (
	"time"

	"github.com/google/uuid"
)

type ContactPerson struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ClientID  uuid.UUID
	Name      string
	Email     *string
	Phone     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
