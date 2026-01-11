package command

import "github.com/google/uuid"

type DeleteBoardScopedCommand struct {
	ID      uuid.UUID
	BoardID uuid.UUID
}
