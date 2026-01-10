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

type InvoiceHandler struct {
	createUC ports.CreateInvoiceUseCase
	updateUC ports.UpdateInvoiceUseCase
	deleteUC ports.DeleteInvoiceUseCase
	findUC   ports.FindInvoiceByIDUseCase
	listUC   ports.ListInvoicesUseCase
	parser   *go_kit.Parser
}

func NewInvoiceHandler(
	createUC ports.CreateInvoiceUseCase,
	updateUC ports.UpdateInvoiceUseCase,
	deleteUC ports.DeleteInvoiceUseCase,
	findUC ports.FindInvoiceByIDUseCase,
	listUC ports.ListInvoicesUseCase,
	parser *go_kit.Parser,
) *InvoiceHandler {
	return &InvoiceHandler{createUC, updateUC, deleteUC, findUC, listUC, parser}
}

type InvoiceListingAttributes struct {
	ExternalID    string `rsql:"filter"`
	ClientID      string `rsql:"filter"`
	IsCanceled    string `rsql:"filter"`
	BillableHours string `rsql:"filter,sort"`
	Total         string `rsql:"filter,sort"`
	IssuedAt      string `rsql:"filter,sort"`
	DueAt         string `rsql:"filter,sort"`
	PaidAt        string `rsql:"filter,sort"`
	Name          string `rsql:"filter,sort"`
	CreatedAt     string `rsql:"filter,sort"`
}

func (h *InvoiceHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &InvoiceListingAttributes{})
	req := &dto.ListInvoicesReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	items, total, err := h.listUC.Execute(c.Request.Context(), req.UserID, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}
	handle.Page(c, listingQuery, total, converter.ToInvoiceListResponse(items))
}

func (h *InvoiceHandler) FindByID(c *gin.Context) {
	req := &dto.FindInvoiceByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	item, err := h.findUC.Execute(c.Request.Context(), req.UserID, req.InvoiceID)
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToInvoiceResponse(item))
}

func (h *InvoiceHandler) Create(c *gin.Context) {
	req := &dto.CreateInvoiceReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	created, err := h.createUC.Execute(c.Request.Context(), converter.ToInvoicePortFromCreate(req))
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusCreated, converter.ToInvoiceResponse(created))
}

func (h *InvoiceHandler) Update(c *gin.Context) {
	req := &dto.UpdateInvoiceReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	updated, err := h.updateUC.Execute(c.Request.Context(), converter.ToInvoicePortFromUpdate(req))
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToInvoiceResponse(updated))
}

func (h *InvoiceHandler) Delete(c *gin.Context) {
	req := &dto.FindInvoiceByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	if err := h.deleteUC.Execute(c.Request.Context(), req.UserID, req.InvoiceID); err != nil {
		handle.Error(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
