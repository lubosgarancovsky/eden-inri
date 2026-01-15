package model

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"

	"time"
)

type Invoice struct {
	ID            uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID        uuid.UUID      `gorm:"type:uuid;not null"`
	ClientID      uuid.UUID      `gorm:"type:uuid;not null"`
	Name          string         `gorm:"type:text;not null"`
	InternalID    string         `gorm:"type:text;not null;unique"`
	Description   *string        `gorm:"type:text"`
	ExternalID    *string        `gorm:"type:text"`
	ExternalLink  *string        `gorm:"type:text"`
	Total         float64        `gorm:"type:double;not null"`
	BillableHours float64        `gorm:"type:double;not null"`
	IsCanceled    bool           `gorm:"type:boolean;not null"`
	IssuedAt      time.Time      `gorm:"type:timestamptz;not null"`
	DueAt         time.Time      `gorm:"type:timestamptz;not null"`
	DeliveredAt   time.Time      `gorm:"type:timestamptz;not null"`
	PaidAt        *time.Time     `gorm:"type:timestamptz"`
	CreatedAt     time.Time      `gorm:"type:timestamptz;not null"`
	UpdatedAt     time.Time      `gorm:"type:timestamptz;not null"`
	DeletedAt     gorm.DeletedAt `gorm:"type:timestamptz;index"`
	Client        Client         `gorm:"foreignKey:ClientID;references:ID;->"`
}

func (Invoice) TableName() string { return "inri_invoice" }

func (m Invoice) ToDomain() *entity.Invoice {
	inv := &entity.Invoice{
		ID:            m.ID,
		UserID:        m.UserID,
		ClientID:      m.ClientID,
		Name:          m.Name,
		Description:   m.Description,
		ExternalID:    m.ExternalID,
		InternalID:    m.InternalID,
		ExternalLink:  m.ExternalLink,
		Total:         m.Total,
		BillableHours: m.BillableHours,
		IsCanceled:    m.IsCanceled,
		IssuedAt:      m.IssuedAt,
		DueAt:         m.DueAt,
		DeliveredAt:   m.DeliveredAt,
		PaidAt:        m.PaidAt,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		Client:        m.Client.ToDomain(),
	}
	return inv
}
