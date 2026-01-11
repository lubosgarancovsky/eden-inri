package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateProjectReq struct {
	UserID      string   `header:"X-User-ID"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Status      string   `json:"status"`
	Tags        []string `json:"tags"`
	Slug        string   `json:"slug"`
}

type UpdateProjectReq struct {
	ID string `uri:"projectId"`
	CreateProjectReq
}

type FavouriteProjectReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
}

type FindProjectByIDReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
}

type DeleteProjectReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
}

type ListProjectsReq struct {
	UserID string `header:"X-User-ID"`
}

type ProjectRes struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	Status         string    `json:"status"`
	Tags           []string  `json:"tags"`
	Slug           string    `json:"slug"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	LastActivityAt time.Time `json:"lastActivityAt"`
	StorySequence  int       `json:"storySequence"`
	Role           string    `json:"role"`
	IsStarred      bool      `json:"isStarred"`
}
