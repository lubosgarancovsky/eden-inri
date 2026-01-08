package domain

import (
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
	ID        uuid.UUID
	StoryID   uuid.UUID
	ActorID   uuid.UUID
	Type      ActivityType
	Payload   string
	CreatedAt time.Time
}
