package command

import "github.com/google/uuid"

type DeleteCommand struct {
	ID     uuid.UUID
	UserID uuid.UUID
}
