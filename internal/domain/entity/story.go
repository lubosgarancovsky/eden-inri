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
	CreatedBy   uuid.UUID
	AssigneeID  *uuid.UUID
	Kind        StoryKind
	Slug        string
	Title       string
	Description string
	Priority    int
	Position    int
	Size        *string
	Estimate    *int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StartDate   *time.Time
	EndDate     *time.Time
}
