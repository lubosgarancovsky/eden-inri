package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"
)

type StoryActivity struct {
	ID        uuid.UUID       `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	StoryID   uuid.UUID       `gorm:"type:uuid;not null"`
	ActorID   uuid.UUID       `gorm:"type:uuid;not null"`
	Type      string          `gorm:"type:string;not null"`
	Payload   json.RawMessage `gorm:"type:jsonb"`
	CreatedAt time.Time       `gorm:"type:timestamptz;autoCreateTime;not null"`
	DeletedAt gorm.DeletedAt  `gorm:"type:timestamptz;index"`

	// Relations
	Actor *User `gorm:"foreignKey:ActorID;references:ID"`
}

func (StoryActivity) TableName() string {
	return "inri_story_activities"
}

func (s StoryActivity) ToDomain() *entity.StoryActivity {
	domainActivity := &entity.StoryActivity{
		ID:        s.ID,
		StoryID:   s.StoryID,
		ActorID:   s.ActorID,
		Type:      entity.ActivityType(s.Type),
		Payload:   s.Payload,
		CreatedAt: s.CreatedAt,
	}

	if s.Actor != nil {
		domainActivity.Actor = *s.Actor.ToDomain()
	}

	return domainActivity
}
