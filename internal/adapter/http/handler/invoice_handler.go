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
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type InvoiceHandler struct {
	createUC  ports.CreateInvoiceUseCase
	updateUC  ports.UpdateInvoiceUseCase
	deleteUC  ports.DeleteInvoiceUseCase
	findUC    ports.FindInvoiceByIDUseCase
	listUC    ports.ListInvoicesUseCase
	analyzeUC ports.AnalyzePdfUseCase
	parser    *go_kit.Parser
}

func NewInvoiceHandler(
	createUC ports.CreateInvoiceUseCase,
	updateUC ports.UpdateInvoiceUseCase,
	deleteUC ports.DeleteInvoiceUseCase,
	findUC ports.FindInvoiceByIDUseCase,
	listUC ports.ListInvoicesUseCase,
	analyzeUC ports.AnalyzePdfUseCase,
	parser *go_kit.Parser,
) *InvoiceHandler {
	return &InvoiceHandler{createUC, updateUC, deleteUC, findUC, listUC, analyzeUC, parser}
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

	query, err := converter.ToListQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	items, total, err := h.listUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}
	invoices := converter.ToListResponse(items, converter.ToInvoiceResponse)
	handle.Page(c, listingQuery, total, invoices)
}

func (h *InvoiceHandler) FindByID(c *gin.Context) {
	req := &dto.FindInvoiceByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findUC.Execute(c.Request.Context(), query)
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

	cmd, err := converter.ToCreateInvoiceCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createUC.Execute(c.Request.Context(), cmd)
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
	cmd, err := converter.ToUpdateInvoiceCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToInvoiceResponse(updated))
}

func (h *InvoiceHandler) Delete(c *gin.Context) {
	req := &dto.DeleteInvoiceReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err = h.deleteUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *InvoiceHandler) ExtractPDF(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		handle.Error(c, go_kit.Wrap(go_kit.ErrBadRequest.WithMessage("invalid multipart form"), err))
		return
	}

	file := form.File["file"]

	if len(file) == 0 {
		handle.Error(c, go_kit.ErrBadRequest.WithMessage("no files provided"))
		return
	}

	pdfFile, err := file[0].Open()
	if err != nil {
		handle.Error(c, err)
		return
	}

	defer pdfFile.Close()

	analysisCmd := &command.FilesCommand{Sources: []*entity.FileSource{
		{
			Name:     file[0].Filename,
			MimeType: file[0].Header.Get("Content-Type"),
			Reader:   pdfFile,
			Size:     file[0].Size,
		},
	}}

	invoice, err := h.analyzeUC.Execute(c.Request.Context(), analysisCmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, invoice)
}
