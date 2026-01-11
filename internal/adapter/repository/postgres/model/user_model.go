package model

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

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

func (u User) ToDomain() *entity.User {
	return &entity.User{
		ID:            u.ID,
		Username:      u.Username,
		Email:         u.Email,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Color:         u.Color,
		AvatarVersion: u.AvatarVersion,
		AvatarMime:    u.AvatarMime,
	}
}
