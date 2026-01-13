package query

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/go-kit"
)

type Query struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

type ScopedQuery struct {
	ID      uuid.UUID
	ScopeID uuid.UUID
	UserID  uuid.UUID
}

type ListQuery struct {
	UserID uuid.UUID
	*go_kit.ListingQuery
}

type ScopedListQuery struct {
	ScopeID uuid.UUID
	UserID  uuid.UUID
	*go_kit.ListingQuery
}
