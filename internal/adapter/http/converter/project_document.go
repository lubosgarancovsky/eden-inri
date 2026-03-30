package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateProjectDocumentCommand(input *dto.CreateProjectDocumentReq) (*command.CreateProjectDocumentCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateProjectDocumentCommand{
		UserID:    userID,
		ProjectID: projectID,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}, nil
}

func ToUpdateProjectDocumentCommand(input *dto.UpdateProjectDocumentReq) (*command.UpdateProjectDocumentCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	id, err := uuid.Parse(input.DocumentID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateProjectDocumentCommand{
		UserID:    userID,
		ProjectID: projectID,
		ID:        id,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}, nil
}

func ToDeleteProjectDocumentCommand(input *dto.DeleteProjectDocumentReq) (*command.DeleteProjectDocumentCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	id, err := uuid.Parse(input.DocumentID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.DeleteProjectDocumentCommand{
		UserID:    userID,
		ID:        id,
		ProjectID: projectID,
	}, nil
}

func ToFindProjectDocumentByIDQuery(input *dto.FindProjectDocumentByIDReq) (*query.FindProjectDocumentByIDQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	documentID, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.FindProjectDocumentByIDQuery{
		UserID:     userID,
		ProjectID:  projectID,
		DocumentID: documentID,
	}, nil
}

func ToListProjectDocumentsQuery(input *dto.ListProjectDocumentsReq, lq *go_kit.ListingQuery) (*query.ListProjectDocumentsQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.ListProjectDocumentsQuery{
		UserID:       userID,
		ProjectID:    projectID,
		ListingQuery: lq,
	}, nil
}

func ToProjectDocumentResponse(e *entity.ProjectDocument) *dto.ProjectDocumentRes {
	return &dto.ProjectDocumentRes{
		ID:        e.ID,
		ProjectID: e.ProjectID,
		Name:      e.Name,
		Content:   e.Content,
		Tags:      e.Tags,
		CreatedBy: ToUserResponse(e.CreatedBy),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
