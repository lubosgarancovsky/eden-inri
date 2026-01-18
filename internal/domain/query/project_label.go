package query

import (
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type ListProjectLabelsQuery struct {
	ProjectID    uuid.UUID
	UserID       uuid.UUID
	ListingQuery *go_kit.ListingQuery
}

type FindProjectLabelByIDQuery struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	ID        uuid.UUID
}
