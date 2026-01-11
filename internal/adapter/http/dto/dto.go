package dto

import "github.com/google/uuid"

type BaseDto interface {
	GetID() uuid.UUID
	GetUserID() uuid.UUID
}

type UserIDDto interface {
	GetUserID() uuid.UUID
}

type UserIDDtoStruct struct {
	UserID string
}
