package handler

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/listing"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

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
	lq, apiErr := helpers.CreateListingQuery(c, h.parser, listing.InvoiceFilter, listing.InvoiceSort)
	if apiErr != nil {
		c.Error(apiErr)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.FindAll(user.ID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
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
	UID, err := helpers.ExtractID(c, "invoiceId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.FindByID(user.ID, UID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
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
	var input model.InvoiceRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Create(user.ID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, result)
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
	var input model.InvoiceRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	UID, err := helpers.ExtractID(c, "invoiceId")
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Update(user.ID, UID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
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
	UID, err := helpers.ExtractID(c, "invoiceId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	_, err = h.s.Delete(user.ID, UID)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
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
	invoiceID, err := helpers.ExtractID(c, "invoiceId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

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

	if err := h.s.SaveAttachments(c, user.ID, invoiceID, files); err != nil {
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
	invoiceID, err := helpers.ExtractID(c, "invoiceId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.ListAttachments(user.ID, invoiceID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}
