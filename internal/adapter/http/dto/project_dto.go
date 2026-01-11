package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateProjectReq struct {
	UserID      uuid.UUID `header:"X-User-ID" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Tags        []string  `json:"tags"`
	Slug        string    `json:"slug"`
}

type UpdateProjectReq struct {
	ID uuid.UUID `uri:"projectId" binding:"required"`
	CreateProjectReq
}

type ProjectRes struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
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

type FavouriteProjectReq struct {
	UserID    uuid.UUID `header:"X-User-ID"`
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
}

func (r FavouriteProjectReq) GetID() uuid.UUID {
	return r.ProjectID
}

func (r FavouriteProjectReq) GetUserID() uuid.UUID {
	return r.UserID
}

type FindProjectByIDReq struct {
	UserID    uuid.UUID `header:"X-User-ID"`
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
}

func (r FindProjectByIDReq) GetID() uuid.UUID {
	return r.ProjectID
}

func (r FindProjectByIDReq) GetUserID() uuid.UUID {
	return r.UserID
}

type DeleteProjectReq struct {
	UserID    uuid.UUID `header:"X-User-ID"`
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
}

func (r DeleteProjectReq) GetID() uuid.UUID {
	return r.ProjectID
}

func (r DeleteProjectReq) GetUserID() uuid.UUID {
	return r.UserID
}

type ListProjectsReq struct {
	UserID uuid.UUID `header:"X-User-ID" binding:"required"`
}

func (r ListProjectsReq) GetUserID() uuid.UUID {
	return r.UserID
}
