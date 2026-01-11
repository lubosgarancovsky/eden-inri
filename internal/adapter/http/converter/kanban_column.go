package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateKanbanColumnCommand(input *dto.CreateKanbanColumnReq) *command.CreateKanbanColumnCommand {
	return &command.CreateKanbanColumnCommand{
		BoardID:  input.BoardID,
		Key:      input.Key,
		Name:     input.Name,
		Type:     input.Type,
		Color:    input.Color,
		Position: input.Position,
	}
}

func ToUpdateKanbanColumnCommand(input *dto.UpdateKanbanColumnReq) *command.UpdateKanbanColumnCommand {
	return &command.UpdateKanbanColumnCommand{
		ID:       input.ColumnID,
		BoardID:  input.BoardID,
		Key:      input.Key,
		Name:     input.Name,
		Type:     input.Type,
		Color:    input.Color,
		Position: input.Position,
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

func ToDeleteBoardScopedCommand(input interface{}) (*command.DeleteBoardScopedCommand, error) {
	boardDto, ok := input.(interface {
		GetBoardID() uuid.UUID
		GetID() uuid.UUID
	})
	if !ok {
		return nil, go_kit.ErrInternalServer
	}
	return &command.DeleteBoardScopedCommand{
		ID:      boardDto.GetID(),
		BoardID: boardDto.GetBoardID(),
	}, nil
}

func ToFindByIDBoardScopedQuery(input interface{}) (*query.FindByIDBoardScopedQuery, error) {
	boardDto, ok := input.(interface {
		GetBoardID() uuid.UUID
		GetID() uuid.UUID
	})
	if !ok {
		return nil, go_kit.ErrInternalServer
	}
	return &query.FindByIDBoardScopedQuery{
		ID:      boardDto.GetID(),
		BoardID: boardDto.GetBoardID(),
	}, nil
}
