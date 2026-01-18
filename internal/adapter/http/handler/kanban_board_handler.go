package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/go-kit"
)

type KanbanBoardHandler struct {
	createUC   ports.CreateKanbanBoardUseCase
	updateUC   ports.UpdateKanbanBoardUseCase
	deleteUC   ports.DeleteKanbanBoardUseCase
	findByIDUC ports.FindKanbanBoardByIDUseCase
	listUC     ports.ListKanbanBoardsUseCase
	parser     *go_kit.Parser
}

func NewKanbanBoardHandler(
	createUC ports.CreateKanbanBoardUseCase,
	updateUC ports.UpdateKanbanBoardUseCase,
	deleteUC ports.DeleteKanbanBoardUseCase,
	findByIDUC ports.FindKanbanBoardByIDUseCase,
	listUC ports.ListKanbanBoardsUseCase,
	parser *go_kit.Parser,
) *KanbanBoardHandler {
	return &KanbanBoardHandler{
		createUC,
		updateUC,
		deleteUC,
		findByIDUC,
		listUC,
		parser,
	}
}

type KanbanBoardListingAttributes struct {
	Name      string `rsql:"filter,sort"`
	Status    string `rsql:"filter,sort"`
	CreatedAt string `rsql:"filter,sort"`
}

func (h *KanbanBoardHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &KanbanBoardListingAttributes{})

	req := &dto.ListKanbanBoardsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	items, total, err := h.listUC.Execute(c.Request.Context(), converter.ToListKanbanBoardQuery(req, listingQuery))
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(items, converter.ToKanbanBoardResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *KanbanBoardHandler) FindByID(c *gin.Context) {
	req := &dto.FindKanbanBoardByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findByIDUC.Execute(c.Request.Context(), converter.ToFindKanbanBoardByIDQuery(req))
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToKanbanBoardResponse(item))
}

func (h *KanbanBoardHandler) Create(c *gin.Context) {
	req := &dto.CreateKanbanBoardReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createUC.Execute(c.Request.Context(), converter.ToCreateKanbanBoardCommand(req))
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusCreated, converter.ToKanbanBoardResponse(created))
}

func (h *KanbanBoardHandler) Update(c *gin.Context) {
	req := &dto.UpdateKanbanBoardReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateUC.Execute(c.Request.Context(), converter.ToUpdateKanbanBoardCommand(req))
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToKanbanBoardResponse(updated))
}

func (h *KanbanBoardHandler) Delete(c *gin.Context) {
	req := &dto.DeleteKanbanBoardReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.deleteUC.Execute(c.Request.Context(), converter.ToDeleteKanbanBoardCommand(req)); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
