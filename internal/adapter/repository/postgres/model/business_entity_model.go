package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"
)

type BusinessEntity struct {
	ID        uuid.UUID      `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null"`
	ICO       string         `gorm:"type:text;not null"`
	DIC       string         `gorm:"type:text;not null"`
	ICDPH     *string        `gorm:"column:ic_dph;type:text"`
	Title     string         `gorm:"type:text;not null"`
	Email     string         `gorm:"type:text;not null"`
	Phone     string         `gorm:"type:text;not null"`
	Street    string         `gorm:"type:text;not null"`
	ZipCode   string         `gorm:"type:text;not null"`
	City      string         `gorm:"type:text;not null"`
	Country   string         `gorm:"type:text;not null"`
	IBAN      string         `gorm:"type:text;not null"`
	SWIFT     string         `gorm:"type:text;not null"`
	CreatedAt time.Time      `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"type:timestamptz;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz;index"`
}

func (BusinessEntity) TableName() string { return "inri_business_entity" }

func (be BusinessEntity) ToDomain() *entity.BusinessEntity {
	return &entity.BusinessEntity{
		ID:        be.ID,
		UserID:    be.UserID,
		ICO:       be.ICO,
		DIC:       be.DIC,
		ICDPH:     be.ICDPH,
		Title:     be.Title,
		Email:     be.Email,
		Phone:     be.Phone,
		Street:    be.Street,
		ZipCode:   be.ZipCode,
		City:      be.City,
		Country:   be.Country,
		IBAN:      be.IBAN,
		SWIFT:     be.SWIFT,
		CreatedAt: be.CreatedAt,
		UpdatedAt: be.UpdatedAt,
	}
}
