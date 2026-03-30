package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

type KanbanColumnHandler struct {
	createUC ports.CreateKanbanColumnUseCase
	updateUC ports.UpdateKanbanColumnUseCase
	deleteUC ports.DeleteKanbanColumnUseCase
	listUC   ports.ListKanbanColumnsUseCase
}

func NewKanbanColumnHandler(
	createUC ports.CreateKanbanColumnUseCase,
	updateUC ports.UpdateKanbanColumnUseCase,
	deleteUC ports.DeleteKanbanColumnUseCase,
	listUC ports.ListKanbanColumnsUseCase,
) *KanbanColumnHandler {
	return &KanbanColumnHandler{
		createUC,
		updateUC,
		deleteUC,
		listUC,
	}
}

func (h *KanbanColumnHandler) List(c *gin.Context) {
	req := &dto.ListKanbanColumnsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListKanbanColumnQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	items, err := h.listUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := make([]*dto.KanbanColumnRes, len(items))
	for i := range items {
		responseItems[i] = converter.ToKanbanColumnResponse(&items[i])
	}

	c.JSON(http.StatusOK, responseItems)
}

func (h *KanbanColumnHandler) Create(c *gin.Context) {
	req := &dto.CreateKanbanColumnReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToCreateKanbanColumnCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusCreated, converter.ToKanbanColumnResponse(created))
}

func (h *KanbanColumnHandler) Update(c *gin.Context) {
	req := &dto.UpdateKanbanColumnReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToUpdateKanbanColumnCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToKanbanColumnResponse(updated))
}

func (h *KanbanColumnHandler) Delete(c *gin.Context) {
	req := &dto.DeleteKanbanColumnReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToDeleteKanbanColumnCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.deleteUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
