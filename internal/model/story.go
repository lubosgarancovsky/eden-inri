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

type Story struct {
	ID          uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	ProjectID   uuid.UUID      `json:"projectId"`
	BoardID     *uuid.UUID     `json:"boardId"`
	ColumnID    *uuid.UUID     `json:"columnId"`
	Slug        string         `json:"slug"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Kind        StoryKind      `json:"kind"`
	AssigneeID  *uuid.UUID     `json:"assigneeId"`
	Priority    int            `json:"priority"`
	Size        *int           `json:"size"`
	Estimate    *time.Duration `json:"estimate" swaggertype:"integer"`
	StartDate   *time.Time     `json:"startDate"`
	EndDate     *time.Time     `json:"endDate"`
	Position    int            `json:"position"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

type StoryRequest struct {
	BoardID     *uuid.UUID     `json:"boardId"`
	ColumnID    *uuid.UUID     `json:"columnId"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Kind        StoryKind      `json:"kind"`
	AssigneeID  *uuid.UUID     `json:"assigneeId"`
	Priority    int            `json:"priority"`
	Size        *int           `json:"size"`
	Estimate    *time.Duration `json:"estimate" swaggertype:"integer"`
	StartDate   *time.Time     `json:"startDate"`
	EndDate     *time.Time     `json:"endDate"`
	Position    int            `json:"position"`
}

func (Story) TableName() string {
	return "inri_stories"
}
