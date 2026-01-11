package dto

import "github.com/google/uuid"

type BaseDto struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

type UserIDDto struct {
	UserID uuid.UUID
}
