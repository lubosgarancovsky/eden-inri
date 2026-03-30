package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type BusinessEntityHandler struct {
	createUC ports.CreateBusinessEntityUseCase
	updateUC ports.UpdateBusinessEntityUseCase
	deleteUC ports.DeleteBusinessEntityUseCase
	findUC   ports.FindBusinessEntityByIDUseCase
	listUC   ports.ListBusinessEntitiesUseCase
	parser   *go_kit.Parser
}

func NewBusinessEntityHandler(
	createUC ports.CreateBusinessEntityUseCase,
	updateUC ports.UpdateBusinessEntityUseCase,
	deleteUC ports.DeleteBusinessEntityUseCase,
	findUC ports.FindBusinessEntityByIDUseCase,
	listUC ports.ListBusinessEntitiesUseCase,
	parser *go_kit.Parser,
) *BusinessEntityHandler {
	return &BusinessEntityHandler{createUC, updateUC, deleteUC, findUC, listUC, parser}
}

type BusinessEntityListingAttributes struct {
	ICO       string `rsql:"filter,sort"`
	DIC       string `rsql:"filter,sort"`
	Title     string `rsql:"filter,sort"`
	Email     string `rsql:"filter,sort"`
	City      string `rsql:"filter,sort"`
	Country   string `rsql:"filter,sort"`
	CreatedAt string `rsql:"filter,sort"`
}

func (h *BusinessEntityHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &BusinessEntityListingAttributes{})
	req := &dto.ListBusinessEntitiesReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListBusinessEntitiesQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	items, total, err := h.listUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	entities := converter.ToListResponse(items, converter.ToBusinessEntityResponse)
	handle.Page(c, listingQuery, total, entities)
}

func (h *BusinessEntityHandler) FindByID(c *gin.Context) {
	req := &dto.FindBusinessEntityByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToBusinessEntityQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToBusinessEntityResponse(item))
}

func (h *BusinessEntityHandler) Create(c *gin.Context) {
	req := &dto.CreateBusinessEntityReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToCreateBusinessEntityCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusCreated, converter.ToBusinessEntityResponse(created))
}

func (h *BusinessEntityHandler) Update(c *gin.Context) {
	req := &dto.UpdateBusinessEntityReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToUpdateBusinessEntityCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToBusinessEntityResponse(updated))
}

func (h *BusinessEntityHandler) Delete(c *gin.Context) {
	req := &dto.DeleteBusinessEntityReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToDeleteBusinessEntityCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	deleted, err := h.deleteUC.Execute(c.Request.Context(), &command.DeleteBusinessEntityCommand{
		ID:     cmd.ID,
		UserID: cmd.UserID,
	})
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToBusinessEntityResponse(deleted))
}
