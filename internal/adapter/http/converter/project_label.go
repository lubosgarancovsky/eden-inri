package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateProjectLabelCommand(input *dto.CreateProjectLabelReq) (*command.CreateProjectLabelCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateProjectLabelCommand{
		UserID:      userID,
		ProjectID:   projectID,
		Name:        input.Name,
		Description: input.Description,
		Color:       input.Color,
	}, nil
}

func ToUpdateProjectLabelCommand(input *dto.UpdateProjectLabelReq) (*command.UpdateProjectLabelCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	id, err := uuid.Parse(input.LabelID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateProjectLabelCommand{
		UserID:      userID,
		ProjectID:   projectID,
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Color:       input.Color,
	}, nil
}

func ToDeleteProjectLabelCommand(input *dto.DeleteProjectLabelReq) (*command.DeleteProjectLabelCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	id, err := uuid.Parse(input.LabelID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.DeleteProjectLabelCommand{
		UserID:    userID,
		ProjectID: projectID,
		ID:        id,
	}, nil
}

func ToFindProjectLabelByIDQuery(input *dto.FindProjectLabelByIDReq) (*query.FindProjectLabelByIDQuery, error) {
	id, err := uuid.Parse(input.LabelID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.FindProjectLabelByIDQuery{
		ID:        id,
		ProjectID: projectID,
		UserID:    userID,
	}, nil
}

func ToListProjectLabelsQuery(input *dto.ListProjectLabelsReq, lq *go_kit.ListingQuery) (*query.ListProjectLabelsQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.ListProjectLabelsQuery{
		UserID:       userID,
		ProjectID:    projectID,
		ListingQuery: lq,
	}, nil
}

func ToProjectLabelResponse(e *entity.ProjectLabel) *dto.ProjectLabelRes {
	return &dto.ProjectLabelRes{
		ID:          e.ID,
		ProjectID:   e.ProjectID,
		Name:        e.Name,
		Description: e.Description,
		Color:       e.Color,
		CreatedAt:   e.CreatedAt,
	}
}
