package entity

import (
	"time"

	"github.com/google/uuid"
)

type BusinessEntity struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ICO       string
	DIC       string
	ICDPH     *string
	Title     string
	Email     string
	Phone     string
	Street    string
	ZipCode   string
	City      string
	Country   string
	IBAN      string
	SWIFT     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
