package model

import "github.com/google/uuid"

type User struct {
	ID            uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Username      string    `gorm:"type:text;unique;not null"`
	Email         string    `gorm:"type:text;unique;not null"`
	FirstName     string    `gorm:"type:text;not null"`
	LastName      string    `gorm:"type:text;not null"`
	Color         string    `gorm:"type:text;not null"`
	AvatarVersion *int      `gorm:"type:text"`
	AvatarMime    *string   `gorm:"type:text"`
}

func (User) TableName() string {
	return "iam_users"
}
