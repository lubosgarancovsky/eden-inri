package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
)

func ToCreateProjectCommand(input *dto.CreateProjectReq) (*command.CreateProjectCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateProjectCommand{
		UserID:        userID,
		Name:          input.Name,
		Description:   input.Description,
		Status:        input.Status,
		Tags:          input.Tags,
		Slug:          input.Slug,
		RepositoryURL: input.RepositoryURL,
	}, nil
}

func ToUpdateProjectCommand(input *dto.UpdateProjectReq) (*command.UpdateProjectCommand, error) {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateProjectCommand{
		ID:            id,
		UserID:        userID,
		Name:          input.Name,
		Description:   input.Description,
		Status:        input.Status,
		Tags:          input.Tags,
		RepositoryURL: input.RepositoryURL,
	}, nil
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
