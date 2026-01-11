package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/attachment"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/go-kit"
)

type AttachmentHandler struct {
	listUC         ports.ListAttachmentsUseCase
	findByIDUC     ports.FindAttachmentByIDUseCase
	deleteUC       ports.DeleteAttachmentUseCase
	utilityService *attachment.AttachmentUtilityService
	parser         *go_kit.Parser
}

func NewAttachmentHandler(
	listUC ports.ListAttachmentsUseCase,
	findByIDUC ports.FindAttachmentByIDUseCase,
	deleteUC ports.DeleteAttachmentUseCase,
	utilityService *attachment.AttachmentUtilityService,
	parser *go_kit.Parser,
) *AttachmentHandler {
	return &AttachmentHandler{
		listUC,
		findByIDUC,
		deleteUC,
		utilityService,
		parser,
	}
}

type AttachmentListingAttributes struct {
	Model        string `rsql:"filter"`
	ModelID      string `rsql:"filter"`
	OriginalName string `rsql:"filter,sort"`
	MimeType     string `rsql:"filter,sort"`
	Size         string `rsql:"filter,sort"`
	CreatedAt    string `rsql:"filter,sort"`
	UpdatedAt    string `rsql:"filter,sort"`
}

func (h *AttachmentHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &AttachmentListingAttributes{})

	req := &dto.ListAttachmentsReq{}
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

	responseItems := converter.ToListResponse(items, converter.ToAttachmentResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *AttachmentHandler) FindByID(c *gin.Context) {
	req := &dto.FindAttachmentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToFindByIDQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findByIDUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToAttachmentResponse(item))
}

func (h *AttachmentHandler) Delete(c *gin.Context) {
	req := &dto.DeleteAttachmentReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.deleteUC.Execute(c.Request.Context(), req.UserID, req.AttachmentID); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AttachmentHandler) Download(c *gin.Context) {
	req := &dto.FindAttachmentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToFindByIDQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findByIDUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	path := h.utilityService.GetFilePath(item)
	c.FileAttachment(path, item.OriginalName)
}
