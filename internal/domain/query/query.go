package query

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/go-kit"
)

type FindByIDQuery struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

type ListQuery struct {
	UserID uuid.UUID
	*go_kit.ListingQuery
}
