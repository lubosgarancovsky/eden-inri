package query

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/go-kit"
)

type ListProjectQuery struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	*go_kit.ListingQuery
}

type FindByIDProjectQuery struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	UserID    uuid.UUID
}
