package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
)

type ProjectRes struct {
	ID             uuid.UUID         `json:"id"`
	Slug           string            `json:"slug"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	Status         string            `json:"status"`
	Tags           []string          `json:"tags" swaggertype:"array,string"`
	StorySequence  int               `json:"storySequence"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
	LastActivityAt time.Time         `json:"lastActivityAt"`
	Role           model.ProjectRole `json:"role" gorm:"->"`      // From ProjectUser //TODO: Remove model after refactoring
	IsStarred      bool              `json:"isStarred" gorm:"->"` // From ProjectUser
}

type CreateProjectReq struct {
	Name        string   `json:"name" binding:"required,min=3,max=20"`
	Description string   `json:"description"`
	Status      string   `json:"status" binding:"required"`
	Tags        []string `json:"tags"`
	Slug        string   `json:"slug" binding:"required,min=2,max=5"`
}

type UpdateProjectReq struct {
	ProjectID   uuid.UUID `uri:"projectId" binding:"required,uuid"`
	Name        string    `json:"name" binding:"required,min=3,max=20"`
	Description string    `json:"description"`
	Status      string    `json:"status" binding:"required"`
	Tags        []string  `json:"tags"`
}
