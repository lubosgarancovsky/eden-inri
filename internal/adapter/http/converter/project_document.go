package converter

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ToCreateProjectDocumentCommand(input *dto.CreateProjectDocumentReq) *command.CreateProjectDocumentCommand {
	return &command.CreateProjectDocumentCommand{
		ProjectID: input.ProjectID,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}
}

func ToUpdateProjectDocumentCommand(input *dto.UpdateProjectDocumentReq) *command.UpdateProjectDocumentCommand {
	return &command.UpdateProjectDocumentCommand{
		ID:        input.DocumentID,
		ProjectID: input.ProjectID,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
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
