package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ToCreateProjectCommand(input *dto.CreateProjectReq) *command.CreateProjectCommand {
	return &command.CreateProjectCommand{
		UserID:        uuid.MustParse(input.UserID),
		Name:          input.Name,
		Description:   input.Description,
		Status:        input.Status,
		Tags:          input.Tags,
		Slug:          input.Slug,
		RepositoryURL: input.RepositoryURL,
	}
}

func ToUpdateProjectCommand(input *dto.UpdateProjectReq) *command.UpdateProjectCommand {
	return &command.UpdateProjectCommand{
		ID:            uuid.MustParse(input.ID),
		UserID:        uuid.MustParse(input.UserID),
		Name:          input.Name,
		Description:   input.Description,
		Status:        input.Status,
		Tags:          input.Tags,
		RepositoryURL: input.RepositoryURL,
	}
}

func ToProjectResponse(e *entity.Project) *dto.ProjectRes {
	return &dto.ProjectRes{
		ID:             e.ID,
		Name:           e.Name,
		Description:    e.Description,
		Status:         e.Status,
		Tags:           e.Tags,
		Slug:           e.Slug,
		RepositoryURL:  e.RepositoryUrl,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
		LastActivityAt: e.LastActivityAt,
		StorySequence:  e.StorySequence,
		Role:           string(e.Role),
		IsStarred:      e.IsStarred,
	}
}
