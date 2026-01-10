package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Color         string    `json:"color"`
	AvatarVersion *int      `json:"avatarVersion"`
	AvatarMime    *string   `json:"avatarMime"`
}

func (User) TableName() string {
	return "iam_users"
}
