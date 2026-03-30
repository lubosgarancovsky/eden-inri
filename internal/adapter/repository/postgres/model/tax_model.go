package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"
)

type Tax struct {
	ID          uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null"`
	Category    string         `gorm:"type:string;not null"`
	Amount      float64        `gorm:"type:numeric;not null"`
	Currency    string         `gorm:"type:text;not null"`
	Description *string        `gorm:"type:text"`
	Reference   *string        `gorm:"type:text"`
	Period      *time.Time     `gorm:"type:date;not null"`
	PaidAt      *time.Time     `gorm:"type:timestamptz"`
	CreatedAt   time.Time      `gorm:"type:timestamptz;not null"`
	UpdatedAt   time.Time      `gorm:"type:timestamptz;not null"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamptz;index"`
}

func (Tax) TableName() string { return "inri_taxes" }

func (t Tax) ToDomain() *entity.Tax {
	return &entity.Tax{
		ID:          t.ID,
		UserID:      t.UserID,
		Category:    entity.TaxCategory(t.Category),
		Amount:      t.Amount,
		Currency:    t.Currency,
		Description: t.Description,
		Reference:   t.Reference,
		Period:      t.Period,
		PaidAt:      t.PaidAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
