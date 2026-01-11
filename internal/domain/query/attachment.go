package query

import (
	"github.com/google/uuid"
)

type FindAttachmentsByModelQuery struct {
	UserID    uuid.UUID
	ModelID   uuid.UUID
	ModelName string
}
