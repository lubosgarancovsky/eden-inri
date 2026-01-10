package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
)

type ProjectRes struct {
	ID             uuid.UUID          `json:"id"`
	Slug           string             `json:"slug"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	Status         string             `json:"status"`
	Tags           []string           `json:"tags"`
	StorySequence  int                `json:"storySequence"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
	LastActivityAt time.Time          `json:"lastActivityAt"`
	Role           models.ProjectRole `json:"role" gorm:"->"`      // From ProjectUser //TODO: Remove model after refactoring
	IsStarred      bool               `json:"isStarred" gorm:"->"` // From ProjectUser
}

type CreateProjectReq struct {
	UserID      uuid.UUID `header:"X-User-ID" binding:"required,uuid"`
	Name        string    `json:"name" binding:"required,min=3,max=20"`
	Description string    `json:"description"`
	Status      string    `json:"status" binding:"required"`
	Tags        []string  `json:"tags"`
	Slug        string    `json:"slug" binding:"required,min=2,max=5"`
}

type UpdateProjectReq struct {
	UserID      uuid.UUID `header:"X-User-ID" binding:"required,uuid"`
	ProjectID   uuid.UUID `uri:"projectId" binding:"required,uuid"`
	Name        string    `json:"name" binding:"required,min=3,max=20"`
	Description string    `json:"description"`
	Status      string    `json:"status" binding:"required"`
	Tags        []string  `json:"tags"`
}

type DeleteProjectReq struct {
	UserID    uuid.UUID `header:"X-User-ID" binding:"required,uuid"`
	ProjectID uuid.UUID `uri:"projectId" binding:"required,uuid"`
}

type FindProjectByIDReq struct {
	UserID    uuid.UUID `header:"X-User-ID" binding:"required,uuid"`
	ProjectID uuid.UUID `uri:"projectId" binding:"required,uuid"`
}

type ListProjectsReq struct {
	UserID uuid.UUID `header:"X-User-ID" binding:"required,uuid"`
}
