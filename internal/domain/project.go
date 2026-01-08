package domain

import (
	"time"

	"github.com/google/uuid"
)

type ProjectStatus string

var (
	ProjectStatusInProgress = ProjectStatus("in_progress")
	ProjectStatusInPlanning = ProjectStatus("in_planning")
	ProjectStatusDone       = ProjectStatus("done")
	ProjectStatusCanceled   = ProjectStatus("canceled")
)

type Project struct {
	ID             uuid.UUID
	Slug           string
	Name           string
	Description    string
	StorySequence  int
	IsStarred      bool
	Tags           []string
	Status         ProjectStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastActivityAt time.Time
}
