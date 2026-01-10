package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/services"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/eden-inri/pkg/types"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

var AttachmentListConfig = types.ListConfig{
	Filter: map[string]string{
		"model":        "model",
		"modelId":      "model_id",
		"originalName": "original_name",
		"mimeType":     "mime_type",
		"size":         "size",
		"createdAt":    "created_at",
		"updatedAt":    "updated_at",
	},
	Sort: map[string]string{
		"originalName": "original_name",
		"mimeType":     "mime_type",
		"size":         "size",
		"createdAt":    "created_at",
		"updatedAt":    "updated_at",
	},
}

// AttachmentHandler handles attachment endpoints
type AttachmentHandler struct {
	parser            *rsql.Parser
	attachmentService *services.AttachmentService
}

// AttachmentPage response for listing
type AttachmentPage struct {
	Items      []models.Attachment
	Page       int
	PageSize   int
	TotalCount int64
}

func NewAttachmentHandler(parser *rsql.Parser, attachmentService *services.AttachmentService) *AttachmentHandler {
	return &AttachmentHandler{parser, attachmentService}
}

// FindAll @Summary      List attachments
// @Description  Returns a paginated list of all attachments
// @Tags         Attachments
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter    query     string  false  "RSQL filter query"
// @Param        sort      query     string  false  "Sort query"
// @Success      200  {object}   AttachmentPage
// @Router       /v1/inri/attachments [get]
// @security GatewayAuth
func (h *AttachmentHandler) FindAll(c *gin.Context) {
	helpers.HandleList(c, h.parser, AttachmentListConfig, h.attachmentService.FindAll)
}

// FindByID @Summary      Get attachment by ID
// @Description  Returns an attachment by its ID
// @Tags         Attachments
// @Accept       json
// @Produce      json
// @Param        attachmentId   path      string  true  "Attachment ID"
// @Success      200  {object}   models.Attachment
// @Router       /v1/inri/attachments/{attachmentId} [get]
// @security GatewayAuth
func (h *AttachmentHandler) FindByID(c *gin.Context) {
	helpers.HandleFindByID(c, "attachmentId", h.attachmentService.FindByID)
}

// Delete @Summary      Delete an attachment
// @Description  Deletes the attachment.
// @Tags         Attachments
// @Accept       json
// @Produce      json
// @Param        attachmentId   path      string  true  "Attachment ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/attachments/{attachmentId} [delete]
// @security GatewayAuth
func (h *AttachmentHandler) Delete(c *gin.Context) {
	helpers.HandleDelete(c, "attachmentId", h.attachmentService.Delete)
}

// Download @Summary      Download attachment file
// @Description  Downloads the attachment file from disk using the stored path
// @Tags         Attachments
// @Produce      application/octet-stream
// @Param        attachmentId   path      string  true  "Attachment ID"
// @Success      200  {file}  file
// @Router       /v1/inri/attachments/{attachmentId}/download [get]
// @security GatewayAuth
func (h *AttachmentHandler) Download(c *gin.Context) {
	attachmentID := helpers.ExtractID(c, "attachmentId")
	userID := helpers.GetUserContext(c).ID

	att, err := h.attachmentService.FindByID(c.Request.Context(), userID, attachmentID)
	if err != nil {
		c.Error(err)
		return
	}

	path := h.attachmentService.GetFilePath(att)
	c.FileAttachment(path, att.OriginalName)
}
