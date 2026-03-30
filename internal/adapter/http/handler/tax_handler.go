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

type TaxHandler struct {
	createUC ports.CreateTaxUseCase
	updateUC ports.UpdateTaxUseCase
	deleteUC ports.DeleteTaxUseCase
	findUC   ports.FindTaxByIDUseCase
	listUC   ports.ListTaxesUseCase
	parser   *go_kit.Parser
}

func NewTaxHandler(
	createUC ports.CreateTaxUseCase,
	updateUC ports.UpdateTaxUseCase,
	deleteUC ports.DeleteTaxUseCase,
	findUC ports.FindTaxByIDUseCase,
	listUC ports.ListTaxesUseCase,
	parser *go_kit.Parser,
) *TaxHandler {
	return &TaxHandler{createUC, updateUC, deleteUC, findUC, listUC, parser}
}

type TaxListingAttributes struct {
	Category  string `rsql:"filter"`
	Amount    string `rsql:"filter,sort"`
	Currency  string `rsql:"filter"`
	Period    string `rsql:"filter,sort"`
	PaidAt    string `rsql:"filter,sort"`
	CreatedAt string `rsql:"filter,sort"`
}

func (h *TaxHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &TaxListingAttributes{})
	req := &dto.ListTaxesReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListTaxesQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	items, total, err := h.listUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	taxes := converter.ToListResponse(items, converter.ToTaxResponse)
	handle.Page(c, listingQuery, total, taxes)
}

func (h *TaxHandler) FindByID(c *gin.Context) {
	req := &dto.FindTaxByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToTaxQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToTaxResponse(item))
}

func (h *TaxHandler) Create(c *gin.Context) {
	req := &dto.CreateTaxReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToCreateTaxCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusCreated, converter.ToTaxResponse(created))
}

func (h *TaxHandler) Update(c *gin.Context) {
	req := &dto.UpdateTaxReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToUpdateTaxCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToTaxResponse(updated))
}

func (h *TaxHandler) Delete(c *gin.Context) {
	req := &dto.DeleteTaxReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToDeleteTaxCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	deleted, err := h.deleteUC.Execute(c.Request.Context(), &command.DeleteTaxCommand{
		ID:     cmd.ID,
		UserID: cmd.UserID,
	})
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToTaxResponse(deleted))
}
