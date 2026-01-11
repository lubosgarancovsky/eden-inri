package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"
)

type Story struct {
	ID          uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ProjectID   uuid.UUID      `gorm:"type:uuid;not null"`
	ColumnID    uuid.UUID      `gorm:"type:uuid;not null"`
	BoardID     uuid.UUID      `gorm:"type:uuid;not null"`
	Slug        string         `gorm:"type:string;not null"`
	Title       string         `gorm:"type:string;not null"`
	Description string         `gorm:"type:text"`
	Kind        string         `gorm:"type:string;not null"`
	AssigneeID  *uuid.UUID     `gorm:"type:uuid"`
	Priority    int            `gorm:"type:int;not null;default:0"`
	Size        *string        `gorm:"type:string"`
	Estimate    *int           `gorm:"type:int;not null;default:0"`
	Position    int            `gorm:"type:int;not null;default:0"`
	CreatedBy   uuid.UUID      `gorm:"type:uuid;not null"`
	CreatedAt   time.Time      `gorm:"type:timestamptz;autoCreateTime;not null"`
	UpdatedAt   time.Time      `gorm:"type:timestamptz;autoUpdateTime;not null"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamptz;index"`
	StartDate   *time.Time     `gorm:"type:timestamptz"`
	EndDate     *time.Time     `gorm:"type:timestamptz"`

	// Relations
	Assignee *User `gorm:"foreignKey:AssigneeID;references:ID"`
	Creator  *User `gorm:"foreignKey:CreatedBy;references:ID"`
}

func (Story) TableName() string {
	return "inri_stories"
}

func (s Story) ToDomain() *entity.Story {
	domainStory := &entity.Story{
		ID:          s.ID,
		ProjectID:   s.ProjectID,
		ColumnID:    s.ColumnID,
		BoardID:     s.BoardID,
		Slug:        s.Slug,
		Title:       s.Title,
		Description: s.Description,
		Kind:        entity.StoryKind(s.Kind),
		AssigneeID:  s.AssigneeID,
		Priority:    s.Priority,
		Size:        s.Size,
		Estimate:    s.Estimate,
		Position:    s.Position,
		CreatedBy:   s.CreatedBy,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
		StartDate:   s.StartDate,
		EndDate:     s.EndDate,
	}

	if s.Assignee != nil {
		domainStory.Assignee = s.Assignee.ToDomain()
	}

	if s.Creator != nil {
		domainStory.Creator = s.Creator.ToDomain()
	}

	return domainStory
}
