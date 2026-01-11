package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ProjectFromDomain(e *entity.Project) *model.Project {
	return &model.Project{
		ID:             e.ID,
		Name:           e.Name,
		Description:    e.Description,
		Status:         e.Status,
		Tags:           e.Tags,
		Slug:           e.Slug,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
		LastActivityAt: e.LastActivityAt,
		StorySequence:  e.StorySequence,
	}
}

func ProjectUserFromDomain(e *entity.ProjectUser) *model.ProjectUser {
	return &model.ProjectUser{
		ProjectID: e.ProjectID,
		UserID:    e.UserID,
		Role:      string(e.Role),
		IsStarred: e.IsStarred,
		JoinedAt:  e.JoinedAt,
	}
}
