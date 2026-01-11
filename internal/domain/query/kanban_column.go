package query

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/go-kit"
)

type ListKanbanColumnsQuery struct {
	BoardID      uuid.UUID
	ProjectID    uuid.UUID
	UserID       uuid.UUID
	ListingQuery *go_kit.ListingQuery
}
