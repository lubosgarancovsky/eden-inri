package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

// AttachmentHandler handles attachment endpoints
type AttachmentHandler struct {
	s      *service.AttachmentService
	parser *rsql.Parser
}

// AttachmentPage response for listing
type AttachmentPage struct {
	Items      []model.Attachment
	Page       int
	PageSize   int
	TotalCount int64
}

func NewAttachmentHandler(parser *rsql.Parser, s *service.AttachmentService) *AttachmentHandler {
	return &AttachmentHandler{s: s, parser: parser}
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
func (h *AttachmentHandler) FindAll(c *gin.Context) {
	lq, apiErr := helpers.CreateListingQuery(c, h.parser, map[string]string{}, map[string]string{})
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

// FindByID @Summary      Get attachment by ID
// @Description  Returns an attachment by its ID
// @Tags         Attachments
// @Accept       json
// @Produce      json
// @Param        attachmentId   path      string  true  "Attachment ID"
// @Success      200  {object}   model.Attachment
// @Router       /v1/inri/attachments/{attachmentId} [get]
func (h *AttachmentHandler) FindByID(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "attachmentId")
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

// Delete @Summary      Delete an attachment
// @Description  Deletes the attachment.
// @Tags         Attachments
// @Accept       json
// @Produce      json
// @Param        attachmentId   path      string  true  "Attachment ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/attachments/{attachmentId} [delete]
func (h *AttachmentHandler) Delete(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "attachmentId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	if _, err := h.s.Delete(user.ID, UID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

// Download @Summary      Download attachment file
// @Description  Downloads the attachment file from disk using the stored path
// @Tags         Attachments
// @Produce      application/octet-stream
// @Param        attachmentId   path      string  true  "Attachment ID"
// @Success      200  {file}  file
// @Router       /v1/inri/attachments/{attachmentId}/download [get]
func (h *AttachmentHandler) Download(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "attachmentId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	att, err := h.s.FindByID(user.ID, UID)
	if err != nil {
		c.Error(err)
		return
	}

	path := h.s.GetFilePath(att)
	c.FileAttachment(path, att.OriginalName)
}
