package query

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/go-kit"
)

type ListStoriesQuery struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	*go_kit.ListingQuery
}

type ListAssignedStoriesQuery struct {
	UserID uuid.UUID
	*go_kit.ListingQuery
}

type FindStoryBySlugQuery struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	Slug      string
}

type FindStoryByIDQuery struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	ID        uuid.UUID
}

type ListStoryActivitiesQuery struct {
	StoryID uuid.UUID
	UserID  uuid.UUID
	*go_kit.ListingQuery
}
