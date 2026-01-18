package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateProjectReq struct {
	UserID        string   `header:"X-User-ID"`
	Name          string   `json:"name"`
	Description   *string  `json:"description"`
	Status        string   `json:"status"`
	Tags          []string `json:"tags"`
	Slug          string   `json:"slug"`
	RepositoryURL *string  `json:"repositoryUrl"`
}

type UpdateProjectReq struct {
	UserID        string   `header:"X-User-ID"`
	ID            string   `uri:"projectId"`
	Name          string   `json:"name"`
	Description   *string  `json:"description"`
	Status        string   `json:"status"`
	Tags          []string `json:"tags"`
	RepositoryURL *string  `json:"repositoryUrl"`
}

type FavouriteProjectReq struct {
	UserID string `header:"X-User-ID"`
	ID     string `uri:"projectId"`
}

type FindProjectByIDReq struct {
	UserID string `header:"X-User-ID"`
	ID     string `uri:"projectId"`
}

type DeleteProjectReq struct {
	UserID string `header:"X-User-ID"`
	ID     string `uri:"projectId"`
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
	RepositoryURL  *string   `json:"repositoryUrl"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	LastActivityAt time.Time `json:"lastActivityAt"`
	StorySequence  int       `json:"storySequence"`
	Role           string    `json:"role"`
	IsStarred      bool      `json:"isStarred"`
}
