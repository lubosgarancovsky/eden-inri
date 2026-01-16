package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

func ToCreateKanbanColumnCommand(input *dto.CreateKanbanColumnReq) *command.CreateKanbanColumnCommand {
	return &command.CreateKanbanColumnCommand{
		UserID:    uuid.MustParse(input.UserID),
		BoardID:   uuid.MustParse(input.BoardID),
		ProjectID: uuid.MustParse(input.ProjectID),
		Key:       input.Key,
		Name:      input.Name,
		Type:      input.Type,
		Color:     input.Color,
		Position:  input.Position,
	}
}

func ToUpdateKanbanColumnCommand(input *dto.UpdateKanbanColumnReq) *command.UpdateKanbanColumnCommand {
	return &command.UpdateKanbanColumnCommand{
		UserID:    uuid.MustParse(input.UserID),
		ID:        uuid.MustParse(input.ColumnID),
		BoardID:   uuid.MustParse(input.BoardID),
		ProjectID: uuid.MustParse(input.ProjectID),
		Key:       input.Key,
		Name:      input.Name,
		Type:      input.Type,
		Color:     input.Color,
		Position:  input.Position,
	}
}

func ToDeleteKanbanColumnCommand(input *dto.DeleteKanbanColumnReq) *command.DeleteKanbanColumnCommand {
	return &command.DeleteKanbanColumnCommand{
		ID:        uuid.MustParse(input.ColumnID),
		BoardID:   uuid.MustParse(input.BoardID),
		ProjectID: uuid.MustParse(input.ProjectID),
		UserID:    uuid.MustParse(input.UserID),
	}
}

func ToListKanbanColumnQuery(input *dto.ListKanbanColumnsReq) *query.ListKanbanColumnsQuery {
	return &query.ListKanbanColumnsQuery{
		BoardID:   uuid.MustParse(input.BoardID),
		UserID:    uuid.MustParse(input.UserID),
		ProjectID: uuid.MustParse(input.ProjectID),
	}
}

func ToKanbanColumnResponse(e *entity.KanbanColumn) *dto.KanbanColumnRes {
	return &dto.KanbanColumnRes{
		ID:        e.ID,
		BoardID:   e.BoardID,
		Key:       e.Key,
		Name:      e.Name,
		Type:      string(e.Type),
		Color:     e.Color,
		Position:  e.Position,
		CreatedAt: e.CreatedAt,
	}
}
