package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateProjectDocumentCommand(input *dto.CreateProjectDocumentReq) *command.CreateProjectDocumentCommand {
	return &command.CreateProjectDocumentCommand{
		UserID:    uuid.MustParse(input.UserID),
		ProjectID: uuid.MustParse(input.ProjectID),
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}
}

func ToUpdateProjectDocumentCommand(input *dto.UpdateProjectDocumentReq) *command.UpdateProjectDocumentCommand {
	return &command.UpdateProjectDocumentCommand{
		UserID:    uuid.MustParse(input.UserID),
		ProjectID: uuid.MustParse(input.ProjectID),
		ID:        uuid.MustParse(input.DocumentID),
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}
}

func ToDeleteProjectDocumentCommand(input *dto.DeleteProjectDocumentReq) *command.DeleteProjectDocumentCommand {
	return &command.DeleteProjectDocumentCommand{
		UserID:    uuid.MustParse(input.UserID),
		ID:        uuid.MustParse(input.DocumentID),
		ProjectID: uuid.MustParse(input.ProjectID),
	}
}

func ToFindProjectDocumentByIDQuery(input *dto.FindProjectDocumentByIDReq) *query.FindProjectDocumentByIDQuery {
	return &query.FindProjectDocumentByIDQuery{
		UserID:     uuid.MustParse(input.UserID),
		ProjectID:  uuid.MustParse(input.ProjectID),
		DocumentID: uuid.MustParse(input.ID),
	}
}

func ToListProjectDocumentsQuery(input *dto.ListProjectDocumentsReq, lq *go_kit.ListingQuery) *query.ListProjectDocumentsQuery {
	return &query.ListProjectDocumentsQuery{
		UserID:       uuid.MustParse(input.UserID),
		ProjectID:    uuid.MustParse(input.ProjectID),
		ListingQuery: lq,
	}
}

func ToProjectDocumentResponse(e *entity.ProjectDocument) *dto.ProjectDocumentRes {
	return &dto.ProjectDocumentRes{
		ID:        e.ID,
		ProjectID: e.ProjectID,
		Name:      e.Name,
		Content:   e.Content,
		Tags:      e.Tags,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
