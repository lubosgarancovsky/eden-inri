package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateContactPersonCommand struct {
	UserID   uuid.UUID
	ClientID uuid.UUID
	Name     string
	Email    *string
	Phone    *string
}

func (c *CreateContactPersonCommand) ToDomain() *entity.ContactPerson {
	return &entity.ContactPerson{
		ID:        uuid.New(),
		UserID:    c.UserID,
		ClientID:  c.ClientID,
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type UpdateContactPersonCommand struct {
	ID uuid.UUID
	CreateContactPersonCommand
}

func (c *UpdateContactPersonCommand) Apply(e *entity.ContactPerson) {
	e.Name = c.Name
	e.Email = c.Email
	e.Phone = c.Phone
	e.UpdatedAt = time.Now()
}

type DeleteContactPersonCommand struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	ClientID uuid.UUID
}
