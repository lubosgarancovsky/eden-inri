package entity

import (
	"time"

	"github.com/google/uuid"
)

type StoryTimeLog struct {
	ID          uuid.UUID
	StoryID     uuid.UUID
	UserID      uuid.UUID
	Duration    time.Duration
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
