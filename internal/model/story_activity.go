package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ActivityType string

const (
	ActivityComment        ActivityType = "comment"
	ActivityChangeColumn   ActivityType = "change_column"
	ActivityAddLabel       ActivityType = "add_label"
	ActivityRemoveLabel    ActivityType = "remove_label"
	ActivityEstimateChange ActivityType = "estimate_change"
	ActivityChangeAssignee ActivityType = "change_assignee"
)

type StoryActivity struct {
	ID        uuid.UUID       `json:"id"`
	StoryID   uuid.UUID       `json:"storyId"`
	ActorID   uuid.UUID       `json:"actorId"`
	Type      ActivityType    `json:"type"`
	Payload   json.RawMessage `json:"payload" swaggertype:"object"`
	CreatedAt time.Time       `json:"createdAt"`
}

type StoryActivityRequest struct {
	Type    ActivityType    `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func (StoryActivity) TableName() string {
	return "inri_story_activities"
}
