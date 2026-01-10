package dto

import (
	"time"

	"github.com/google/uuid"
)

type StoryActivity struct {
	ID        uuid.UUID `json:"id"`
	StoryID   uuid.UUID `json:"storyId"`
	Actor     UserRes   `json:"actor"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"` // TODO: This is JSON, so I could sent it parsed
	CreatedAt time.Time `json:"createdAt"`
}

type CreateStoryActivityReq struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

type UpdateStoryActivityReq struct {
	StoryID uuid.UUID `uri:"storyId"`
	Payload string    `json:"payload"`
}
