package model

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"time"
)

type Label struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ProjectID   uuid.UUID
	Name        string
	Description string
	Color       string
	CreatedAt   time.Time
}

func (Label) TableName() string { return "inri_labels" }

func (m *Label) ToPort() *ports.Label {
	return &ports.Label{ID: m.ID, ProjectID: m.ProjectID, Name: m.Name, Description: m.Description, Color: m.Color}
}

func LabelFromPort(p *ports.Label) *Label {
	return &Label{ID: p.ID, ProjectID: p.ProjectID, Name: p.Name, Description: p.Description, Color: p.Color}
}
