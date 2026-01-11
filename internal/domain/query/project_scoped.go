package query

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/go-kit"
)

type ListProjectScopedQuery struct {
	ProjectID uuid.UUID
	*go_kit.ListingQuery
}

type FindByIDProjectScopedQuery struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
}
