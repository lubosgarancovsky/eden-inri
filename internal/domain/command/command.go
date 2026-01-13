package command

import "github.com/google/uuid"

type Command struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

type ScopedCommand struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	ScopeID uuid.UUID
}
