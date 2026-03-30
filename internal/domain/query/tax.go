package query

import (
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type ListTaxesQuery struct {
	UserID uuid.UUID
	*go_kit.ListingQuery
}

type FindTaxByIDQuery struct {
	ID     uuid.UUID
	UserID uuid.UUID
}
