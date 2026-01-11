package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func KanbanColumnFromDomain(e *entity.KanbanColumn) *model.KanbanColumn {
	return &model.KanbanColumn{
		ID:       e.ID,
		BoardID:  e.BoardID,
		Key:      e.Key,
		Name:     e.Name,
		Type:     string(e.Type),
		Color:    e.Color,
		Position: e.Position,
	}
}
