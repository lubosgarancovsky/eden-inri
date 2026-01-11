package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"
)

type Client struct {
	ID           uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null"`
	ClientType   string         `gorm:"type:string;not null"`
	ContractType string         `gorm:"type:string;not null"`
	Name         string         `gorm:"type:string;not null"`
	Description  *string        `gorm:"type:string"`
	TaxNumber    *string        `gorm:"type:string"`
	Address      *string        `gorm:"type:string"`
	Tags         pq.StringArray `gorm:"type:text[];not null"`
	HourRate     float64        `gorm:"type:float;not null"`
	StartedAt    time.Time      `gorm:"type:timestamptz;not null"`
	FinishedAt   *time.Time     `gorm:"type:timestamptz"`
	CreatedAt    time.Time      `gorm:"type:timestamptz;autoCreateTime;not null"`
	UpdatedAt    time.Time      `gorm:"type:timestamptz;autoUpdateTime;not null"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamptz;index"`
}

func (Client) TableName() string {
	return "inri_clients"
}

func (c Client) ToDomain() *entity.Client {
	return &entity.Client{
		ID:           c.ID,
		UserID:       c.UserID,
		ClientType:   entity.ClientType(c.ClientType),
		ContractType: entity.ClientContractType(c.ContractType),
		Name:         c.Name,
		Description:  c.Description,
		TaxNumber:    c.TaxNumber,
		Address:      c.Address,
		Tags:         c.Tags,
		HourRate:     c.HourRate,
		StartedAt:    c.StartedAt,
		FinishedAt:   c.FinishedAt,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}
