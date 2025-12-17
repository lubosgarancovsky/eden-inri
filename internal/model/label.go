package model

import (
	"time"

	"github.com/google/uuid"
)

type Label struct {
	ID          uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	ProjectID   *uuid.UUID `json:"projectId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Color       string     `json:"color"`
	CreatedAt   time.Time  `json:"createdAt"`
}

type LabelRequest struct {
	ProjectID   *uuid.UUID `json:"projectId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Color       string     `json:"color"`
}

type StoryLabel struct {
	StoryID uuid.UUID `json:"storyId"`
	LabelID uuid.UUID `json:"labelId"`
}

func (StoryLabel) TableName() string {
	return "inri_story_labels"
}

func (Label) TableName() string {
	return "inri_labels"
}
