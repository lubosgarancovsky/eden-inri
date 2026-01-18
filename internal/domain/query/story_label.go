package query

import "github.com/google/uuid"

type ListStoryLabelsQuery struct {
	UserID    uuid.UUID
	ProjectID uuid.UUID
	StoryID   uuid.UUID
}
