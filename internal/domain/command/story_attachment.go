package command

import "github.com/google/uuid"

type DeleteStoryAttachmentCommand struct {
	UserID       uuid.UUID
	ProjectID    uuid.UUID
	StoryID      uuid.UUID
	AttachmentID uuid.UUID
}
