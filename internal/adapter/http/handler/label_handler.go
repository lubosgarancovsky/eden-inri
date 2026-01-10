package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type LabelHandler struct {
	createUC ports.CreateLabelUseCase
	updateUC ports.UpdateLabelUseCase
	deleteUC ports.DeleteLabelUseCase
	findUC   ports.FindLabelByIDUseCase
	listUC   ports.ListLabelsUseCase
	parser   *go_kit.Parser
}

func NewLabelHandler(createUC ports.CreateLabelUseCase, updateUC ports.UpdateLabelUseCase, deleteUC ports.DeleteLabelUseCase, findUC ports.FindLabelByIDUseCase, listUC ports.ListLabelsUseCase, parser *go_kit.Parser) *LabelHandler {
	return &LabelHandler{createUC, updateUC, deleteUC, findUC, listUC, parser}
}

type LabelListingAttributes struct {
	ID        string `rsql:"filter"`
	ProjectID string `rsql:"filter"`
	BoardID   string `rsql:"filter"`
	Name      string `rsql:"filter,sort"`
	Color     string `rsql:"filter"`
	CreatedAt string `rsql:"filter,sort"`
}

func (h *LabelHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &LabelListingAttributes{})
	req := &dto.ListLabelsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	items, total, err := h.listUC.Execute(c.Request.Context(), req.ProjectID, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}
	handle.Page(c, listingQuery, total, converter.ToLabelListResponse(items))
}

func (h *LabelHandler) FindByID(c *gin.Context) {
	req := &dto.FindLabelByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	item, err := h.findUC.Execute(c.Request.Context(), req.ProjectID, req.LabelID)
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToLabelResponse(item))
}

func (h *LabelHandler) Create(c *gin.Context) {
	req := &dto.CreateLabelReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	created, err := h.createUC.Execute(c.Request.Context(), converter.ToLabelPortFromCreate(req))
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusCreated, converter.ToLabelResponse(created))
}

func (h *LabelHandler) Update(c *gin.Context) {
	req := &dto.UpdateLabelReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	updated, err := h.updateUC.Execute(c.Request.Context(), converter.ToLabelPortFromUpdate(req))
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToLabelResponse(updated))
}

func (h *LabelHandler) Delete(c *gin.Context) {
	req := &dto.FindLabelByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	if err := h.deleteUC.Execute(c.Request.Context(), req.ProjectID, req.LabelID); err != nil {
		handle.Error(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
