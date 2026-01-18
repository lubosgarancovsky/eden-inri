package query

import (
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type FindByIDKanbanBoardQuery struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	UserID    uuid.UUID
}

type ListKanbanBoardQuery struct {
	ProjectID    uuid.UUID
	UserID       uuid.UUID
	ListingQuery *go_kit.ListingQuery
}
