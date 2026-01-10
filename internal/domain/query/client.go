package query

import (
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type FindClientByIDQuery struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

type ListClientsQuery struct {
	UserID uuid.UUID
	go_kit.ListingQuery
}
