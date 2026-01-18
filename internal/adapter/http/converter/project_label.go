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
		UserID:      uuid.MustParse(input.UserID),
		ProjectID:   uuid.MustParse(input.ProjectID),
		Name:        input.Name,
		Description: input.Description,
		Color:       input.Color,
	}
}

func ToUpdateProjectLabelCommand(input *dto.UpdateProjectLabelReq) *command.UpdateProjectLabelCommand {
	return &command.UpdateProjectLabelCommand{
		UserID:      uuid.MustParse(input.UserID),
		ProjectID:   uuid.MustParse(input.ProjectID),
		ID:          uuid.MustParse(input.LabelID),
		Name:        input.Name,
		Description: input.Description,
		Color:       input.Color,
	}
}

func ToDeleteProjectLabelCommand(input *dto.DeleteProjectLabelReq) *command.DeleteProjectLabelCommand {
	return &command.DeleteProjectLabelCommand{
		UserID:    uuid.MustParse(input.UserID),
		ProjectID: uuid.MustParse(input.ProjectID),
		ID:        uuid.MustParse(input.LabelID),
	}
}

func ToFindProjectLabelByIDQuery(input *dto.FindProjectLabelByIDReq) *query.FindProjectLabelByIDQuery {
	return &query.FindProjectLabelByIDQuery{
		ID:        uuid.MustParse(input.LabelID),
		ProjectID: uuid.MustParse(input.ProjectID),
		UserID:    uuid.MustParse(input.UserID),
	}
}

func ToListProjectLabelsQuery(input *dto.ListProjectLabelsReq, lq *go_kit.ListingQuery) *query.ListProjectLabelsQuery {
	return &query.ListProjectLabelsQuery{
		UserID:       uuid.MustParse(input.UserID),
		ProjectID:    uuid.MustParse(input.ProjectID),
		ListingQuery: lq,
	}
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
