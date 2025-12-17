package model

import (
	"time"

	"github.com/google/uuid"
)

type TimeLog struct {
	ID          uuid.UUID     `json:"id"`
	StoryID     uuid.UUID     `json:"storyId"`
	UserID      uuid.UUID     `json:"userId"`
	Duration    time.Duration `json:"duration"`
	Description *string       `json:"description"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

type TimeLogRequest struct {
	Duration    time.Duration `json:"duration"`
	Description *string       `json:"description"`
}

func (TimeLog) TableName() string {
	return "inri_story_time_logs"
}
