package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToListStoryLabelsQuery(input *dto.ListStoryLabelReq) (*query.ListStoryLabelsQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	storyID, err := uuid.Parse(input.StoryID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &query.ListStoryLabelsQuery{
		UserID:    userID,
		ProjectID: projectID,
		StoryID:   storyID,
	}, nil
}

func ToStoryLabelCommand(input *dto.StoryLabelReq) (*command.StoryLabelCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	storyID, err := uuid.Parse(input.StoryID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	labelID, err := uuid.Parse(input.LabelID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &command.StoryLabelCommand{
		UserID:    userID,
		ProjectID: projectID,
		StoryID:   storyID,
		LabelID:   labelID,
	}, nil
}

func ToStoryLabelResponse(input *entity.StoryLabel) *dto.StoryLabelRes {
	return &dto.StoryLabelRes{
		dto.ProjectLabelRes{
			ID:          input.Label.ID,
			ProjectID:   input.Label.ProjectID,
			Name:        input.Label.Name,
			Description: input.Label.Description,
			Color:       input.Label.Color,
			CreatedAt:   input.Label.CreatedAt,
		},
	}
}
