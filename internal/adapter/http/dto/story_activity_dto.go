package dto

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type CreateStoryActivityReq struct {
	UserID  string          `header:"X-User-ID"`
	StoryID string          `uri:"storyId"`
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type UpdateStoryActivityReq struct {
	UserID     string          `header:"X-User-ID"`
	ActivityID string          `uri:"activityId"`
	Payload    json.RawMessage `json:"payload"`
}

type DeleteStoryActivityReq struct {
	UserID     string `header:"X-User-ID"`
	ActivityID string `uri:"activityId"`
}

type ListStoryActivitiesReq struct {
	UserID  string `header:"X-User-ID"`
	StoryID string `uri:"storyId"`
}

type StoryActivityRes struct {
	ID        uuid.UUID       `json:"id"`
	StoryID   uuid.UUID       `json:"storyId"`
	ActorID   uuid.UUID       `json:"actorId"`
	Type      string          `json:"type"`
	Payload   json.RawMessage `json:"payload"`
	CreatedAt time.Time       `json:"createdAt"`
	Actor     *UserRes        `json:"actor,omitempty"`
}
