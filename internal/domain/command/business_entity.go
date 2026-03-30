package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateBusinessEntityCommand struct {
	UserID  uuid.UUID
	ICO     string
	DIC     string
	ICDPH   *string
	Title   string
	Email   string
	Phone   string
	Street  string
	ZipCode string
	City    string
	Country string
	IBAN    string
	SWIFT   string
}

func (c *CreateBusinessEntityCommand) ToDomain() *entity.BusinessEntity {
	now := time.Now()
	return &entity.BusinessEntity{
		ID:        uuid.New(),
		UserID:    c.UserID,
		ICO:       c.ICO,
		DIC:       c.DIC,
		ICDPH:     c.ICDPH,
		Title:     c.Title,
		Email:     c.Email,
		Phone:     c.Phone,
		Street:    c.Street,
		ZipCode:   c.ZipCode,
		City:      c.City,
		Country:   c.Country,
		IBAN:      c.IBAN,
		SWIFT:     c.SWIFT,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type UpdateBusinessEntityCommand struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	ICO     string
	DIC     string
	ICDPH   *string
	Title   string
	Email   string
	Phone   string
	Street  string
	ZipCode string
	City    string
	Country string
	IBAN    string
	SWIFT   string
}

func (c *UpdateBusinessEntityCommand) Apply(e *entity.BusinessEntity) {
	e.ICO = c.ICO
	e.DIC = c.DIC
	e.ICDPH = c.ICDPH
	e.Title = c.Title
	e.Email = c.Email
	e.Phone = c.Phone
	e.Street = c.Street
	e.ZipCode = c.ZipCode
	e.City = c.City
	e.Country = c.Country
	e.IBAN = c.IBAN
	e.SWIFT = c.SWIFT
	e.UpdatedAt = time.Now()
}

type DeleteBusinessEntityCommand struct {
	ID     uuid.UUID
	UserID uuid.UUID
}
