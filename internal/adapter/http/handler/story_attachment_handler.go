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

type StoryAttachmentHandler struct {
	listUC     ports.ListStoryAttachmentsUserCase
	findByIDUC ports.FindStoryAttachmentByIDUseCase
	deleteUC   ports.DeleteStoryAttachmentUseCase
	parser     *go_kit.Parser
}

func NewStoryAttachmentHandler(
	listUC ports.ListStoryAttachmentsUserCase,
	findByIDUC ports.FindStoryAttachmentByIDUseCase,
	deleteUC ports.DeleteStoryAttachmentUseCase,
	parser *go_kit.Parser,
) *StoryAttachmentHandler {
	return &StoryAttachmentHandler{
		listUC,
		findByIDUC,
		deleteUC,
		parser,
	}
}

func (h *StoryAttachmentHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &AttachmentListingAttributes{})

	req := &dto.ListStoryAttachmentsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToScopedListQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	items, totalCount, err := h.listUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(items, converter.ToAttachmentResponse)
	handle.Page(c, listingQuery, totalCount, responseItems)
}

func (h *StoryAttachmentHandler) Delete(c *gin.Context) {
	req := &dto.StoryAttachmentByIDReq{}
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

func (h *StoryAttachmentHandler) Download(c *gin.Context) {
	req := &dto.StoryAttachmentByIDReq{}
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
