package query

import (
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type FindProjectByIDQuery struct {
	UserID    uuid.UUID
	ProjectID uuid.UUID
}

type ListProjectsQuery struct {
	UserID       uuid.UUID
	ListingQuery go_kit.ListingQuery
}
