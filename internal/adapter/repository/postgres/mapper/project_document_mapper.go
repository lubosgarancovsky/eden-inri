package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ProjectDocumentFromDomain(e *entity.ProjectDocument) *model.ProjectDocument {
	return &model.ProjectDocument{
		ID:        e.ID,
		ProjectID: e.ProjectID,
		Name:      e.Name,
		Content:   e.Content,
		Tags:      e.Tags,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
