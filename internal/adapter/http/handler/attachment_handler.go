package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/go-kit"
)

type AttachmentHandler struct {
	listUC     ports.ListAttachmentsUseCase
	findByIDUC ports.FindAttachmentByIDUseCase
	deleteUC   ports.DeleteAttachmentUseCase
	uploadUC   ports.UploadAttachmentUseCase
	updateUC   ports.UpdateAttachmentUseCase
	parser     *go_kit.Parser
}

func NewAttachmentHandler(
	listUC ports.ListAttachmentsUseCase,
	findByIDUC ports.FindAttachmentByIDUseCase,
	deleteUC ports.DeleteAttachmentUseCase,
	uploadUC ports.UploadAttachmentUseCase,
	updateUC ports.UpdateAttachmentUseCase,
	parser *go_kit.Parser,
) *AttachmentHandler {
	return &AttachmentHandler{
		listUC,
		findByIDUC,
		deleteUC,
		uploadUC,
		updateUC,
		parser,
	}
}

type AttachmentListingAttributes struct {
	Model        string `rsql:"filter,sort"`
	ModelID      string `rsql:"field:model_id,filter"`
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
	req := &dto.AttachmentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToQuery(req)
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

func (h *AttachmentHandler) Update(c *gin.Context) {
	req := &dto.UpdateAttachmentReq{}
	if err := validator.BindAndValidate(c, &req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd := converter.ToUpdateAttachmentCommand(req)
	attachment, err := h.updateUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToAttachmentResponse(attachment))
}

func (h *AttachmentHandler) Delete(c *gin.Context) {
	req := &dto.AttachmentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToCommand(req)
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

func (h *AttachmentHandler) Upload(c *gin.Context) {
	userID := handle.UserID(c)
	model := c.PostForm("model")
	modelID := c.PostForm("modelId")

	req := &dto.UploadAttachmentReq{
		UserID:  userID.String(),
		ModelID: modelID,
		Model:   model,
	}

	form, err := c.MultipartForm()
	if err != nil {
		handle.Error(c, go_kit.Wrap(go_kit.ErrBadRequest.WithMessage("invalid multipart form"), err))
		return
	}
	files := form.File["files"]

	if len(files) == 0 {
		handle.Error(c, go_kit.ErrBadRequest.WithMessage("no files provided"))
		return
	}

	if err = h.uploadUC.Execute(c.Request.Context(), converter.ToUploadAttachmentCommand(req, files)); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AttachmentHandler) Download(c *gin.Context) {
	req := &dto.AttachmentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findByIDUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.Header("Content-Disposition", `attachment; filename="`+item.OriginalName+`"`)
	c.Header("Content-Type", item.MimeType)

	fmt.Println(item.GetFilePath(config.GlobalConfig.UploadsFolder))

	c.File(item.GetFilePath(config.GlobalConfig.UploadsFolder))
}
