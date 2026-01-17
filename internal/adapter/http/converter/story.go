package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	"github.com/lubosgarancovsky/go-kit"
)

func ToCreateStoryCommand(input *dto.CreateStoryReq) *command.CreateStoryCommand {
	var assigneeID *uuid.UUID
	if input.AssigneeID != nil {
		id := uuid.MustParse(*input.AssigneeID)
		assigneeID = &id
	}

	return &command.CreateStoryCommand{
		UserID:      uuid.MustParse(input.UserID),
		ProjectID:   uuid.MustParse(input.ProjectID),
		ColumnID:    uuid.MustParse(input.ColumnID),
		BoardID:     uuid.MustParse(input.BoardID),
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
	}
}

func ToUpdateStoryCommand(input *dto.UpdateStoryReq) *command.UpdateStoryCommand {
	var assigneeID *uuid.UUID
	if input.AssigneeID != nil {
		id := uuid.MustParse(*input.AssigneeID)
		assigneeID = &id
	}

	return &command.UpdateStoryCommand{
		ID:        uuid.MustParse(input.ID),
		UserID:    uuid.MustParse(input.UserID),
		ProjectID: uuid.MustParse(input.ProjectID),
		CreateStoryCommand: command.CreateStoryCommand{
			UserID:      uuid.MustParse(input.UserID),
			ProjectID:   uuid.MustParse(input.ProjectID),
			ColumnID:    uuid.MustParse(input.ColumnID),
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
	}
}

func ToDeleteStoryCommand(input *dto.DeleteStoryReq) *command.DeleteStoryCommand {
	return &command.DeleteStoryCommand{
		ID:        uuid.MustParse(input.ID),
		UserID:    uuid.MustParse(input.UserID),
		ProjectID: uuid.MustParse(input.ProjectID),
	}
}

func ToChangeStoryAssigneeCommand(input *dto.ChangeStoryAssigneeReq) *command.ChangeStoryAssigneeCommand {
	var assigneeID *uuid.UUID
	if input.AssigneeID != nil {
		id := uuid.MustParse(*input.AssigneeID)
		assigneeID = &id
	}

	return &command.ChangeStoryAssigneeCommand{
		ID:         uuid.MustParse(input.ID),
		UserID:     uuid.MustParse(input.UserID),
		ProjectID:  uuid.MustParse(input.ProjectID),
		AssigneeID: assigneeID,
	}
}

func ToFindStoryByIDQuery(input *dto.FindStoryByIDReq) *query.FindStoryByIDQuery {
	return &query.FindStoryByIDQuery{
		ID:        uuid.MustParse(input.ID),
		UserID:    uuid.MustParse(input.UserID),
		ProjectID: uuid.MustParse(input.ProjectID),
	}
}

func ToFindStoryBySlugQuery(input *dto.FindStoryBySlugReq) *query.FindStoryBySlugQuery {
	return &query.FindStoryBySlugQuery{
		Slug:      input.Slug,
		UserID:    uuid.MustParse(input.UserID),
		ProjectID: uuid.MustParse(input.ProjectID),
	}
}

func ToListStoriesQuery(input *dto.ListStoriesReq, lq *go_kit.ListingQuery) *query.ListStoriesQuery {
	return &query.ListStoriesQuery{
		UserID:       uuid.MustParse(input.UserID),
		ProjectID:    uuid.MustParse(input.ProjectID),
		ListingQuery: lq,
	}
}

func ToListAssignedStoriesQuery(input *dto.ListAssignedStoriesReq, lq *go_kit.ListingQuery) *query.ListAssignedStoriesQuery {
	return &query.ListAssignedStoriesQuery{
		UserID:       uuid.MustParse(input.UserID),
		ListingQuery: lq,
	}
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
