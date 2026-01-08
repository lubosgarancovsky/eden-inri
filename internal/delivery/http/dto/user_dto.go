package dto

import "github.com/google/uuid"

type UserRes struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Color         string    `json:"color"`
	AvatarVersion *int      `json:"avatarVersion"`
	AvatarMime    *string   `json:"avatarMime"`
}
