package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateProjectLabelCommand(input *dto.CreateProjectLabelReq) *command.CreateProjectLabelCommand {
	return &command.CreateProjectLabelCommand{
		ProjectID:   input.ProjectID,
		Name:        input.Name,
		Description: input.Description,
		Color:       input.Color,
	}
}

func ToUpdateProjectLabelCommand(input *dto.UpdateProjectLabelReq) *command.UpdateProjectLabelCommand {
	return &command.UpdateProjectLabelCommand{
		ID:          input.LabelID,
		ProjectID:   input.ProjectID,
		Name:        input.Name,
		Description: input.Description,
		Color:       input.Color,
	}
}

func ToProjectLabelResponse(e *entity.ProjectLabel) *dto.ProjectLabelRes {
	var projectID uuid.UUID
	if e.ProjectID != nil {
		projectID = *e.ProjectID
	}
	return &dto.ProjectLabelRes{
		ID:          e.ID,
		ProjectID:   projectID,
		Name:        e.Name,
		Description: e.Description,
		Color:       e.Color,
		CreatedAt:   e.CreatedAt,
	}
}

func ToDeleteProjectScopedCommand(input interface{}) (*command.DeleteProjectScopedCommand, error) {
	projectDto, ok := input.(interface {
		GetProjectID() uuid.UUID
		GetID() uuid.UUID
	})
	if !ok {
		return nil, go_kit.ErrInternalServer
	}
	return &command.DeleteProjectScopedCommand{
		ID:        projectDto.GetID(),
		ProjectID: projectDto.GetProjectID(),
	}, nil
}

func ToFindByIDProjectScopedQuery(input interface{}) (*query.FindByIDProjectScopedQuery, error) {
	projectDto, ok := input.(interface {
		GetProjectID() uuid.UUID
		GetID() uuid.UUID
	})
	if !ok {
		return nil, go_kit.ErrInternalServer
	}
	return &query.FindByIDProjectScopedQuery{
		ID:        projectDto.GetID(),
		ProjectID: projectDto.GetProjectID(),
	}, nil
}

func ToListProjectScopedQuery(input interface{}, lq *go_kit.ListingQuery) (*query.ListProjectScopedQuery, error) {
	projectDto, ok := input.(interface{ GetProjectID() uuid.UUID })
	if !ok {
		return nil, go_kit.ErrInternalServer
	}
	return &query.ListProjectScopedQuery{
		ProjectID:    projectDto.GetProjectID(),
		ListingQuery: lq,
	}, nil
}
