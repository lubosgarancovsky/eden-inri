package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

type StoryLabelHandler struct {
	listUC     ports.ListStoryLabelsUseCase
	assignUC   ports.AssignStoryLabelUseCase
	unassignUC ports.UnassignStoryLabelUseCase
}

func NewStoryLabelHandler(listUC ports.ListStoryLabelsUseCase, assignUC ports.AssignStoryLabelUseCase, unassignUC ports.UnassignStoryLabelUseCase) *StoryLabelHandler {
	return &StoryLabelHandler{listUC, assignUC, unassignUC}
}

func (h *StoryLabelHandler) List(c *gin.Context) {
	req := &dto.ListStoryLabelReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListStoryLabelsQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	items, err := h.listUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(items, converter.ToStoryLabelResponse)
	c.JSON(http.StatusOK, responseItems)
}

func (h *StoryLabelHandler) Assign(c *gin.Context) {
	req := &dto.StoryLabelReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToStoryLabelCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.assignUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *StoryLabelHandler) Unassign(c *gin.Context) {
	req := &dto.StoryLabelReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToStoryLabelCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.unassignUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
