package model

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

type StoryListItem struct {
	ID         uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	ProjectID  uuid.UUID  `json:"projectId"`
	ColumnID   uuid.UUID  `json:"columnId"`
	Slug       string     `json:"slug"`
	Title      string     `json:"title"`
	Kind       StoryKind  `json:"kind"`
	AssigneeID *uuid.UUID `json:"-"`
	Priority   int        `json:"priority" gorm:"not null;default:0"`
	Size       *string    `json:"size"`
	Estimate   *int       `json:"estimate" gorm:"not null;default:0"`
	Position   int        `json:"position" gorm:"not null;default:0"`
	Assignee   *User      `json:"assignee" gorm:"column:assignee_id"`
}

type Story struct {
	StoryListItem
	Description string     `json:"description"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"<-:create"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	CreatedBy   uuid.UUID  `json:"-" gorm:"column:created_by;<-:create"`
	Creator     *User      `json:"creator" gorm:"foreignKey:CreatedBy;references:ID"`
}

type StoryRequest struct {
	ColumnID    uuid.UUID  `json:"columnId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Kind        StoryKind  `json:"kind"`
	AssigneeID  *uuid.UUID `json:"assigneeId"`
	Priority    int        `json:"priority"`
	Size        *string    `json:"size"`
	Estimate    *int       `json:"estimate"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
	Position    int        `json:"position"`
}

type StoryLabelRequest struct {
	LabelID uuid.UUID `json:"labelId"`
}

func (Story) TableName() string {
	return "inri_stories"
}
