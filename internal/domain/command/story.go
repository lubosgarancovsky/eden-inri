package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateStoryCommand struct {
	UserID      uuid.UUID
	ProjectID   uuid.UUID
	ColumnID    uuid.UUID
	Title       string
	Description string
	Kind        string
	AssigneeID  *uuid.UUID
	Priority    int
	Size        *string
	Estimate    *int
	StartDate   *time.Time
	EndDate     *time.Time
	Position    int
}

func (c *CreateStoryCommand) ToDomain() *entity.Story {
	return &entity.Story{
		ID:          uuid.New(),
		ProjectID:   c.ProjectID,
		ColumnID:    c.ColumnID,
		Title:       c.Title,
		Description: c.Description,
		Kind:        entity.StoryKind(c.Kind),
		AssigneeID:  c.AssigneeID,
		Priority:    c.Priority,
		Size:        c.Size,
		Estimate:    c.Estimate,
		StartDate:   c.StartDate,
		EndDate:     c.EndDate,
		Position:    c.Position,
		CreatedBy:   c.UserID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type UpdateStoryCommand struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProjectID uuid.UUID
	CreateStoryCommand
}

func (c *UpdateStoryCommand) Apply(e *entity.Story) {
	e.ColumnID = c.ColumnID
	e.Title = c.Title
	e.Description = c.Description
	e.Kind = entity.StoryKind(c.Kind)
	e.AssigneeID = c.AssigneeID
	e.Priority = c.Priority
	e.Size = c.Size
	e.Estimate = c.Estimate
	e.StartDate = c.StartDate
	e.EndDate = c.EndDate
	e.Position = c.Position
	e.UpdatedAt = time.Now()
}

type ChangeStoryAssigneeCommand struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	ProjectID  uuid.UUID
	AssigneeID *uuid.UUID
}

type DeleteStoryCommand struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProjectID uuid.UUID
}
