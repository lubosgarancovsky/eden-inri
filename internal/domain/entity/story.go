package entity

import (
	"time"

	"github.com/google/uuid"
)

type StoryKind string

const (
	StoryBug      StoryKind = "bug"
	StoryFeature  StoryKind = "feature"
	StoryDoc      StoryKind = "doc"
	StoryTask     StoryKind = "task"
	StoryDesign   StoryKind = "design"
	StoryPlanning StoryKind = "planning"
)

type Story struct {
	ID          uuid.UUID
	ProjectID   uuid.UUID
	ColumnID    uuid.UUID
	BoardID     uuid.UUID
	Slug        string
	Title       string
	Description string
	Kind        StoryKind
	AssigneeID  *uuid.UUID
	Priority    int
	Size        *string
	Estimate    *int
	Position    int
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StartDate   *time.Time
	EndDate     *time.Time

	// Relations (optional in domain, but useful if they are returned by repo)
	Assignee *User
	Creator  *User
}
