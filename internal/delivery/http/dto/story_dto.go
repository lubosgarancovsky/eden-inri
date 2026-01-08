package dto

import (
	"time"

	"github.com/google/uuid"
)

type StoryRes struct {
	ID        uuid.UUID  `json:"id"`
	ProjectID uuid.UUID  `json:"projectId"`
	ColumnID  uuid.UUID  `json:"columnId"`
	BoardID   uuid.UUID  `json:"boardId"`
	Slug      string     `json:"slug"`
	Title     string     `json:"title"`
	Kind      string     `json:"kind"`
	Priority  int        `json:"priority"`
	Size      *string    `json:"size"`
	Estimate  *int       `json:"estimate"`
	Position  int        `json:"position"`
	Assignee  *UserRes   `json:"assignee"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}

type StoryDetailRes struct {
	StoryRes
	Description string   `json:"description"`
	Creator     *UserRes `json:"creator"`
}

type CreateStoryReq struct {
	ColumnID    uuid.UUID  `json:"columnId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Kind        string     `json:"kind"`
	AssigneeID  *uuid.UUID `json:"assigneeId"`
	Priority    int        `json:"priority"`
	Size        *string    `json:"size"`
	Estimate    *int       `json:"estimate"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
	Position    int        `json:"position"`
}
