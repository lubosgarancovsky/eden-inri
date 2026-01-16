package query

import (
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type ListStoryAttachmentsQuery struct {
	UserID       uuid.UUID
	ProjectID    uuid.UUID
	StoryID      uuid.UUID
	ListingQuery *go_kit.ListingQuery
}

type FindStoryAttachmentByIDQuery struct {
	UserID       uuid.UUID
	ProjectID    uuid.UUID
	StoryID      uuid.UUID
	AttachmentID uuid.UUID
}
