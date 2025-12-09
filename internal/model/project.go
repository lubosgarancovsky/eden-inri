package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Project struct {
	ID          uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID      uuid.UUID      `json:"-"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Status      string         `json:"status"`
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags" swaggertype:"array,string"`
	Slug        string         `json:"slug"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

type ProjectRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	Tags        []string `json:"tags"`
	Slug        string   `json:"slug"`
}

func (p *Project) TableName() string {
	return "inri_projects"
}
