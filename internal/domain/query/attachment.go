package query

import (
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type FindAttachmentsByModelQuery struct {
	UserID    uuid.UUID
	ModelID   uuid.UUID
	ModelName string
}

type ListAttachmentsByModelQuery struct {
	UserID       uuid.UUID
	ModelID      uuid.UUID
	ModelName    string
	ListingQuery *go_kit.ListingQuery
}
