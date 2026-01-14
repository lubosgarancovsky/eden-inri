package entity

import "github.com/google/uuid"

type StoryLabel struct {
	StoryID uuid.UUID
	Label   *ProjectLabel
}
