package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateKanbanColumnReq struct {
	BoardID  uuid.UUID `uri:"boardId" binding:"required"`
	Key      string    `json:"key" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Type     string    `json:"type" binding:"required"`
	Color    string    `json:"color"`
	Position int       `json:"position" binding:"required"`
}

type UpdateKanbanColumnReq struct {
	BoardID  uuid.UUID `uri:"boardId" binding:"required"`
	ColumnID uuid.UUID `uri:"columnId" binding:"required"`
	Key      string    `json:"key" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Type     string    `json:"type" binding:"required"`
	Color    string    `json:"color"`
	Position int       `json:"position" binding:"required"`
}

type DeleteKanbanColumnReq struct {
	BoardID  uuid.UUID `uri:"boardId" binding:"required"`
	ColumnID uuid.UUID `uri:"columnId" binding:"required"`
}

type FindKanbanColumnByIDReq struct {
	BoardID  uuid.UUID `uri:"boardId" binding:"required"`
	ColumnID uuid.UUID `uri:"columnId" binding:"required"`
}

type ListKanbanColumnsReq struct {
	BoardID uuid.UUID `uri:"boardId" binding:"required"`
}

type KanbanColumnRes struct {
	ID        uuid.UUID `json:"id"`
	BoardID   uuid.UUID `json:"boardId"`
	Key       string    `json:"key"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Color     string    `json:"color"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"createdAt"`
}

func (r CreateKanbanColumnReq) GetID() uuid.UUID      { return uuid.Nil }
func (r CreateKanbanColumnReq) GetUserID() uuid.UUID  { return uuid.Nil }
func (r CreateKanbanColumnReq) GetBoardID() uuid.UUID { return r.BoardID }

func (r UpdateKanbanColumnReq) GetID() uuid.UUID      { return r.ColumnID }
func (r UpdateKanbanColumnReq) GetUserID() uuid.UUID  { return uuid.Nil }
func (r UpdateKanbanColumnReq) GetBoardID() uuid.UUID { return r.BoardID }

func (r DeleteKanbanColumnReq) GetID() uuid.UUID      { return r.ColumnID }
func (r DeleteKanbanColumnReq) GetUserID() uuid.UUID  { return uuid.Nil }
func (r DeleteKanbanColumnReq) GetBoardID() uuid.UUID { return r.BoardID }

func (r FindKanbanColumnByIDReq) GetID() uuid.UUID      { return r.ColumnID }
func (r FindKanbanColumnByIDReq) GetUserID() uuid.UUID  { return uuid.Nil }
func (r FindKanbanColumnByIDReq) GetBoardID() uuid.UUID { return r.BoardID }

func (r ListKanbanColumnsReq) GetUserID() uuid.UUID  { return uuid.Nil }
func (r ListKanbanColumnsReq) GetBoardID() uuid.UUID { return r.BoardID }
