package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func StoryFromDomain(e *entity.Story) *model.Story {
	return &model.Story{
		ID:          e.ID,
		ProjectID:   e.ProjectID,
		ColumnID:    e.ColumnID,
		BoardID:     e.BoardID,
		Slug:        e.Slug,
		Title:       e.Title,
		Description: e.Description,
		Kind:        string(e.Kind),
		AssigneeID:  e.AssigneeID,
		Priority:    e.Priority,
		Size:        e.Size,
		Estimate:    e.Estimate,
		Position:    e.Position,
		CreatedBy:   e.CreatedBy,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		StartDate:   e.StartDate,
		EndDate:     e.EndDate,
	}
}
