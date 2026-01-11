package query

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/go-kit"
)

type FindProjectDocumentByIDQuery struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProjectID uuid.UUID
}

type ListProjectDocumentsQuery struct {
	UserID    uuid.UUID
	ProjectID uuid.UUID
	go_kit.ListingQuery
}
