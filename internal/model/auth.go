package model

import "github.com/google/uuid"

type UserContext struct {
	ID   uuid.UUID
	Role string
}
