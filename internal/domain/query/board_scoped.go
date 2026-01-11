package query

import (
	"github.com/google/uuid"
)

type FindByIDBoardScopedQuery struct {
	ID      uuid.UUID
	BoardID uuid.UUID
}
