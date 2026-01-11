package query

import (
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type ListContactPersonsQuery struct {
	UserID   uuid.UUID
	ClientID uuid.UUID
	*go_kit.ListingQuery
}

type FindContactPersonByIDQuery struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	ClientID uuid.UUID
}
