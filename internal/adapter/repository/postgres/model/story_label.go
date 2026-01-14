package model

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type StoryLabel struct {
	StoryID uuid.UUID    `gorm:"type:uuid"`
	LabelID uuid.UUID    `gorm:"type:uuid"`
	Label   ProjectLabel `gorm:"->"`
}

func (StoryLabel) TableName() string {
	return "inri_story_labels"
}

func (sl *StoryLabel) ToDomain() *entity.StoryLabel {
	return &entity.StoryLabel{
		StoryID: sl.StoryID,
		Label:   sl.Label.ToDomain(),
	}
}
