package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type ProjectAttachmentHandler struct {
	listUC     ports.ListProjectAttachmentsUseCase
	findByIDUC ports.FindProjectAttachmentByIDUseCase
	updateUC   ports.UpdateProjectAttachmentUseCase
	deleteUC   ports.DeleteProjectAttachmentUseCase
	parser     *go_kit.Parser
}

func NewProjectAttachmentHandler(
	listUC ports.ListProjectAttachmentsUseCase,
	findByIDUC ports.FindProjectAttachmentByIDUseCase,
	updateUC ports.UpdateProjectAttachmentUseCase,
	deleteUC ports.DeleteProjectAttachmentUseCase,
	parser *go_kit.Parser,
) *ProjectAttachmentHandler {
	return &ProjectAttachmentHandler{
		listUC,
		findByIDUC,
		updateUC,
		deleteUC,
		parser,
	}
}

func (h *ProjectAttachmentHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &AttachmentListingAttributes{})

	req := &dto.ListProjectAttachmentsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToScopedListQuery(req, listingQuery)
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

func (h *ProjectAttachmentHandler) FindByID(c *gin.Context) {
	req := &dto.ProjectAttachmentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToScopedQuery(req)
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

func (h *ProjectAttachmentHandler) Update(c *gin.Context) {
	req := &dto.UpdateProjectAttachmentReq{}
	if err := validator.BindAndValidate(c, &req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToUpdateProjectAttachmentCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	attachment, err := h.updateUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToAttachmentResponse(attachment))
}

func (h *ProjectAttachmentHandler) Delete(c *gin.Context) {
	req := &dto.ProjectAttachmentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToScopedCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.deleteUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ProjectAttachmentHandler) Download(c *gin.Context) {
	req := &dto.ProjectAttachmentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToScopedQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findByIDUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.FileAttachment(item.GetFilePath(config.GlobalConfig.UploadsFolder), item.OriginalName)
}
