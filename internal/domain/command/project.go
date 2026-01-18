package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateProjectCommand struct {
	UserID        uuid.UUID
	Name          string
	Description   *string
	Status        string
	Tags          []string
	Slug          string
	RepositoryURL *string
}

func (c *CreateProjectCommand) ToDomain() *entity.Project {
	return &entity.Project{
		ID:             uuid.New(),
		Name:           c.Name,
		Description:    c.Description,
		Status:         c.Status,
		Tags:           c.Tags,
		Slug:           c.Slug,
		RepositoryUrl:  c.RepositoryURL,
		Role:           entity.ProjectRoleOwner,
		IsStarred:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		LastActivityAt: time.Now(),
	}
}

type UpdateProjectCommand struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	Name          string
	Description   *string
	Status        string
	Tags          []string
	RepositoryURL *string
}

func (c *UpdateProjectCommand) Apply(e *entity.Project) {
	e.Name = c.Name
	e.Description = c.Description
	e.Status = c.Status
	e.Tags = c.Tags
	e.RepositoryUrl = c.RepositoryURL
	e.UpdatedAt = time.Now()
}

type DeleteProjectCommand struct {
	ID     uuid.UUID
	UserID uuid.UUID
}
