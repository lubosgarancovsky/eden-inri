package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateStoryReq struct {
	UserID      string     `header:"X-User-ID"`
	ProjectID   string     `uri:"projectId"`
	ColumnID    string     `json:"columnId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Kind        string     `json:"kind"`
	AssigneeID  *string    `json:"assigneeId"`
	Priority    int        `json:"priority"`
	Size        *string    `json:"size"`
	Estimate    *int       `json:"estimate"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
	Position    int        `json:"position"`
}

type UpdateStoryReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	ID        string `uri:"storyId"`
	CreateStoryReq
}

type DeleteStoryReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	ID        string `uri:"storyId"`
}

type FindStoryByIDReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	ID        string `uri:"storyId"`
}

type ListStoriesReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
}

type ListAssignedStoriesReq struct {
	UserID string `header:"X-User-ID"`
}

type ChangeStoryAssigneeReq struct {
	UserID     string  `header:"X-User-ID"`
	ProjectID  string  `uri:"projectId"`
	ID         string  `uri:"storyId"`
	AssigneeID *string `json:"assigneeId"`
}

type StoryAttachmentReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	StoryID   string `uri:"storyId"`
}

type StoryRes struct {
	ID          uuid.UUID  `json:"id"`
	ProjectID   uuid.UUID  `json:"projectId"`
	ColumnID    uuid.UUID  `json:"columnId"`
	BoardID     uuid.UUID  `json:"boardId"`
	Slug        string     `json:"slug"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Kind        string     `json:"kind"`
	AssigneeID  *uuid.UUID `json:"assigneeId"`
	Priority    int        `json:"priority"`
	Size        *string    `json:"size"`
	Estimate    *int       `json:"estimate"`
	Position    int        `json:"position"`
	CreatedBy   uuid.UUID  `json:"createdBy"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
	Assignee    *UserRes   `json:"assignee,omitempty"`
	Creator     *UserRes   `json:"creator,omitempty"`
}

type UserRes struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Color     string    `json:"color"`
}
