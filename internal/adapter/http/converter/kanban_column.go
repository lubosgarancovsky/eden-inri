package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	"github.com/lubosgarancovsky/go-kit"
)

func ToCreateKanbanColumnCommand(input *dto.CreateKanbanColumnReq) (*command.CreateKanbanColumnCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	boardID, err := uuid.Parse(input.BoardID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateKanbanColumnCommand{
		UserID:    userID,
		BoardID:   boardID,
		ProjectID: projectID,
		Key:       input.Key,
		Name:      input.Name,
		Type:      input.Type,
		Color:     input.Color,
		Position:  input.Position,
	}, nil
}

func ToUpdateKanbanColumnCommand(input *dto.UpdateKanbanColumnReq) (*command.UpdateKanbanColumnCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	id, err := uuid.Parse(input.ColumnID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	boardID, err := uuid.Parse(input.BoardID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateKanbanColumnCommand{
		UserID:    userID,
		ID:        id,
		BoardID:   boardID,
		ProjectID: projectID,
		Key:       input.Key,
		Name:      input.Name,
		Type:      input.Type,
		Color:     input.Color,
		Position:  input.Position,
	}, nil
}

func ToDeleteKanbanColumnCommand(input *dto.DeleteKanbanColumnReq) (*command.DeleteKanbanColumnCommand, error) {
	id, err := uuid.Parse(input.ColumnID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	boardID, err := uuid.Parse(input.BoardID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.DeleteKanbanColumnCommand{
		ID:        id,
		BoardID:   boardID,
		ProjectID: projectID,
		UserID:    userID,
	}, nil
}

func ToListKanbanColumnQuery(input *dto.ListKanbanColumnsReq) (*query.ListKanbanColumnsQuery, error) {
	boardID, err := uuid.Parse(input.BoardID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.ListKanbanColumnsQuery{
		BoardID:   boardID,
		UserID:    userID,
		ProjectID: projectID,
	}, nil
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
