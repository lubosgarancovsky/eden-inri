package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	"github.com/lubosgarancovsky/go-kit"
)

func ToCreateStoryCommand(input *dto.CreateStoryReq) (*command.CreateStoryCommand, error) {
	var assigneeID *uuid.UUID
	if input.AssigneeID != nil {
		id, err := uuid.Parse(*input.AssigneeID)
		if err != nil {
			return nil, go_kit.ErrInvalidUUID
		}
		assigneeID = &id
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	columnID, err := uuid.Parse(input.ColumnID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	boardID, err := uuid.Parse(input.BoardID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &command.CreateStoryCommand{
		UserID:      userID,
		ProjectID:   projectID,
		ColumnID:    columnID,
		BoardID:     boardID,
		Title:       input.Title,
		Description: input.Description,
		Kind:        input.Kind,
		AssigneeID:  assigneeID,
		Priority:    input.Priority,
		Size:        input.Size,
		Estimate:    input.Estimate,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Position:    input.Position,
	}, nil
}

func ToUpdateStoryCommand(input *dto.UpdateStoryReq) (*command.UpdateStoryCommand, error) {
	var assigneeID *uuid.UUID
	if input.AssigneeID != nil {
		id, err := uuid.Parse(*input.AssigneeID)
		if err != nil {
			return nil, go_kit.ErrInvalidUUID
		}
		assigneeID = &id
	}

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	columnID, err := uuid.Parse(input.ColumnID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &command.UpdateStoryCommand{
		ID:        id,
		UserID:    userID,
		ProjectID: projectID,
		CreateStoryCommand: command.CreateStoryCommand{
			UserID:      userID,
			ProjectID:   projectID,
			ColumnID:    columnID,
			Title:       input.Title,
			Description: input.Description,
			Kind:        input.Kind,
			AssigneeID:  assigneeID,
			Priority:    input.Priority,
			Size:        input.Size,
			Estimate:    input.Estimate,
			StartDate:   input.StartDate,
			EndDate:     input.EndDate,
			Position:    input.Position,
		},
	}, nil
}

func ToDeleteStoryCommand(input *dto.DeleteStoryReq) (*command.DeleteStoryCommand, error) {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &command.DeleteStoryCommand{
		ID:        id,
		UserID:    userID,
		ProjectID: projectID,
	}, nil
}

func ToChangeStoryAssigneeCommand(input *dto.ChangeStoryAssigneeReq) (*command.ChangeStoryAssigneeCommand, error) {
	var assigneeID *uuid.UUID
	if input.AssigneeID != nil {
		id, err := uuid.Parse(*input.AssigneeID)
		if err != nil {
			return nil, go_kit.ErrInvalidUUID
		}
		assigneeID = &id
	}

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &command.ChangeStoryAssigneeCommand{
		ID:         id,
		UserID:     userID,
		ProjectID:  projectID,
		AssigneeID: assigneeID,
	}, nil
}

func ToFindStoryByIDQuery(input *dto.FindStoryByIDReq) (*query.FindStoryByIDQuery, error) {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &query.FindStoryByIDQuery{
		ID:        id,
		UserID:    userID,
		ProjectID: projectID,
	}, nil
}

func ToFindStoryBySlugQuery(input *dto.FindStoryBySlugReq) (*query.FindStoryBySlugQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &query.FindStoryBySlugQuery{
		Slug:      input.Slug,
		UserID:    userID,
		ProjectID: projectID,
	}, nil
}

func ToListStoriesQuery(input *dto.ListStoriesReq, lq *go_kit.ListingQuery) (*query.ListStoriesQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &query.ListStoriesQuery{
		UserID:       userID,
		ProjectID:    projectID,
		ListingQuery: lq,
	}, nil
}

func ToListAssignedStoriesQuery(input *dto.ListAssignedStoriesReq, lq *go_kit.ListingQuery) (*query.ListAssignedStoriesQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &query.ListAssignedStoriesQuery{
		UserID:       userID,
		ListingQuery: lq,
	}, nil
}

func ToStoryResponse(e *entity.Story) *dto.StoryRes {
	res := &dto.StoryRes{
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

	if e.Assignee != nil {
		res.Assignee = ToUserResponse(e.Assignee)
	}

	if e.Creator != nil {
		res.Creator = ToUserResponse(e.Creator)
	}

	return res
}
