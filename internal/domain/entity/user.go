package entity

import "github.com/google/uuid"

type User struct {
	ID            uuid.UUID
	Username      string
	Email         string
	FirstName     string
	LastName      string
	Color         string
	AvatarVersion *int
	AvatarMime    *string
}
