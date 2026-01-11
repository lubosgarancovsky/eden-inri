package command

import "github.com/google/uuid"

type DeleteProjectScopedCommand struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
}
