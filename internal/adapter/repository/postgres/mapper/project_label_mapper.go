package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ProjectLabelFromDomain(e *entity.ProjectLabel) *model.ProjectLabel {
	return &model.ProjectLabel{
		ID:          e.ID,
		ProjectID:   e.ProjectID,
		Name:        e.Name,
		Description: e.Description,
		Color:       e.Color,
		CreatedAt:   e.CreatedAt,
	}
}
