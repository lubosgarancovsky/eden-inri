package dto

import (
	"time"

	"github.com/google/uuid"
)

type StoryTimeLogRes struct {
	ID          uuid.UUID     `json:"id"`
	StoryID     uuid.UUID     `json:"storyId"`
	UserID      uuid.UUID     `json:"userId"`
	Duration    time.Duration `json:"duration"`
	Description *string       `json:"description"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

type CreateTimeLogReq struct {
	Duration    time.Duration `json:"duration"`
	Description *string       `json:"description"`
}
