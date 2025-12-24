package handler

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/eden-inri/pkg/types"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

var InvoiceListConfig = types.ListConfig{
	Filter: map[string]string{
		"externalId":    "external_id",
		"clientId":      "client_id",
		"isCanceled":    "is_canceled",
		"billableHours": "billable_hours",
		"total":         "total",
		"issuedAt":      "issued_at",
		"dueAt":         "due_at",
		"paidAt":        "paid_at",
		"name":          "name",
	},
	Sort: map[string]string{
		"issuedAt":      "issued_at",
		"dueAt":         "due_at",
		"createdAt":     "created_at",
		"billableHours": "billable_hours",
		"total":         "total",
		"paidAt":        "paid_at",
		"name":          "name",
	},
}

type InvoiceHandler struct {
	s      *service.InvoiceService
	parser *rsql.Parser
}

type InvoicePage struct {
	Items      []model.Invoice
	Page       int
	PageSize   int
	TotalCount int64
}

func NewInvoiceHandler(parser *rsql.Parser, s *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{s, parser}
}

// FindAll @Summary      List invoices
// @Description  Returns a paginated list of all invoices
// @Tags         Invoices
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}   InvoicePage
// @Router       /v1/inri/invoices [get]
func (h *InvoiceHandler) FindAll(c *gin.Context) {
	helpers.HandleList(c, h.parser, InvoiceListConfig, h.s.FindAll)
}

// FindByID @Summary      Get invoice by ID
// @Description  Returns an invoice by its ID
// @Tags         Invoices
// @Accept       json
// @Produce      json
// @Param        invoiceId   path      string  true  "Invoice ID"
// @Success      200  {object}   model.Invoice
// @Router       /v1/inri/invoices/{invoiceId} [get]
func (h *InvoiceHandler) FindByID(c *gin.Context) {
	helpers.HandleFindByID(c, "invoiceId", h.s.FindByID)
}

// Create @Summary      Create a new invoice
// @Description  Creates a new invoice
// @Tags         Invoices
// @Accept       json
// @Produce      json
// @Param        invoice  body  model.InvoiceRequest  true  "Invoice data"
// @Success      201  {object}  model.Invoice
// @Router       /v1/inri/invoices [post]
func (h *InvoiceHandler) Create(c *gin.Context) {
	helpers.HandleCreate(c, h.s.Create)
}

// Update @Summary      Update an invoice
// @Description  Updates an existing invoice
// @Tags         Invoices
// @Accept       json
// @Produce      json
// @Param        invoice  body  model.InvoiceRequest  true  "Invoice data"
// @Param        invoiceId   path      string  true  "Invoice ID"
// @Success      200  {object}  model.Invoice
// @Router       /v1/inri/invoices/{invoiceId} [put]
func (h *InvoiceHandler) Update(c *gin.Context) {
	helpers.HandleUpdate(c, "invoiceId", h.s.Update)
}

// Delete @Summary      Delete an invoice
// @Description  Deletes the invoice
// @Tags         Invoices
// @Accept       json
// @Produce      json
// @Param        invoiceId   path      string  true  "Invoice ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/invoices/{invoiceId} [delete]
func (h *InvoiceHandler) Delete(c *gin.Context) {
	helpers.HandleDelete(c, "invoiceId", h.s.Delete)
}

// UploadAttachments @Summary Upload attachments to an invoice
// @Description Upload multiple files as attachments for the given invoice
// @Tags         Invoices
// @Accept       mpfd
// @Produce      json
// @Param        invoiceId   path      string  true  "Invoice ID"
// @Param files formData []file true "Files to upload"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/invoices/{invoiceId}/attachments [post]
func (h *InvoiceHandler) UploadAttachments(c *gin.Context) {
	invoiceID := helpers.ExtractID(c, "invoiceId")
	userID := helpers.GetUserContext(c).ID

	if err := c.Request.ParseMultipartForm(32 << 20); err != nil { // 32MB default limit
		c.Error(err)
		return
	}
	form := c.Request.MultipartForm
	var files []multipart.FileHeader
	if fhs, ok := form.File["files"]; ok {
		for _, fh := range fhs {
			files = append(files, *fh)
		}
	}
	if len(files) == 0 {
		// also support single file key "file"
		if f, err2 := c.FormFile("file"); err2 == nil && f != nil {
			files = append(files, *f)
		}
	}

	if len(files) == 0 {
		c.Error(api_err.ErrBadRequest.WithMessage("no files provided"))
		return
	}

	if err := h.s.SaveAttachments(c, userID, invoiceID, files); err != nil {
		c.Error(err)
		return
	}
	c.Status(204)
}

// ListAttachments @Summary      List invoice attachments
// @Description  Returns a list of attachments by invoice ID
// @Tags         Invoices
// @Accept       json
// @Produce      json
// @Success      200  {object}   []model.Attachment
// @Router       /v1/inri/invoices/{invoiceId}/attachments [get]
func (h *InvoiceHandler) ListAttachments(c *gin.Context) {
	invoiceID := helpers.ExtractID(c, "invoiceId")
	userID := helpers.GetUserContext(c).ID

	result, err := h.s.ListAttachments(c.Request.Context(), userID, invoiceID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}
