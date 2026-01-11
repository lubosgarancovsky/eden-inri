package query

import (
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type FindProjectDocumentByIDQuery struct {
	ProjectID  uuid.UUID
	DocumentID uuid.UUID
	UserID     uuid.UUID
}

type ListProjectDocumentsQuery struct {
	ProjectID    uuid.UUID
	UserID       uuid.UUID
	ListingQuery *go_kit.ListingQuery
}
