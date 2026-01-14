package command

import "github.com/google/uuid"

type StoryLabelCommand struct {
	UserID    uuid.UUID
	ProjectID uuid.UUID
	StoryID   uuid.UUID
	LabelID   uuid.UUID
}
