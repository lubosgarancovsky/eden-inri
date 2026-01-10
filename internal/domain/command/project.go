package command

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateProjectCommand struct {
	UserID      uuid.UUID
	Name        string
	Description string
	Status      string
	Tags        []string
	Slug        string
}

func (c *CreateProjectCommand) ToDomain() *entity.Project {
	// User ownership is managed at persistence/query time (via context/userID)
	return &entity.Project{
		ID:          uuid.New(),
		Slug:        c.Slug,
		Name:        c.Name,
		Description: c.Description,
		Status:      entity.ProjectStatus(c.Status),
		Tags:        c.Tags,
	}
}

type UpdateProjectCommand struct {
	UserID      uuid.UUID
	ProjectID   uuid.UUID
	Name        string
	Description string
	Status      string
	Tags        []string
}

func (c *UpdateProjectCommand) Apply(p *entity.Project) {
	p.Name = c.Name
	p.Description = c.Description
	p.Status = entity.ProjectStatus(c.Status)
	p.Tags = c.Tags
}

func (c *UpdateProjectCommand) ToDomain() *entity.Project {
	return &entity.Project{
		ID:          c.ProjectID,
		Name:        c.Name,
		Description: c.Description,
		Status:      entity.ProjectStatus(c.Status),
		Tags:        c.Tags,
	}
}

type DeleteProjectCommand struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
}
