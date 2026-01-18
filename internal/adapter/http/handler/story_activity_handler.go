package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/go-kit"
)

type StoryActivityHandler struct {
	createStoryActivityUC ports.CreateStoryActivityUseCase
	updateStoryActivityUC ports.UpdateStoryActivityUseCase
	deleteStoryActivityUC ports.DeleteStoryActivityUseCase
	listStoryActivitiesUC ports.ListStoryActivitiesUseCase
	parser                *go_kit.Parser
}

func NewStoryActivityHandler(
	createStoryActivityUC ports.CreateStoryActivityUseCase,
	updateStoryActivityUC ports.UpdateStoryActivityUseCase,
	deleteStoryActivityUC ports.DeleteStoryActivityUseCase,
	listStoryActivitiesUC ports.ListStoryActivitiesUseCase,
	parser *go_kit.Parser,
) *StoryActivityHandler {
	return &StoryActivityHandler{
		createStoryActivityUC,
		updateStoryActivityUC,
		deleteStoryActivityUC,
		listStoryActivitiesUC,
		parser,
	}
}

type StoryActivityListingAttributes struct {
	CreatedAt string `rsql:"filter,sort"`
	Type      string `rsql:"filter"`
}

func (h *StoryActivityHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &StoryActivityListingAttributes{})

	req := &dto.ListStoryActivitiesReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query := converter.ToListStoryActivitiesQuery(req, listingQuery)

	activities, total, err := h.listStoryActivitiesUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(activities, converter.ToStoryActivityResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *StoryActivityHandler) Create(c *gin.Context) {
	req := &dto.CreateStoryActivityReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createStoryActivityUC.Execute(c.Request.Context(), converter.ToCreateStoryActivityCommand(req))
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusCreated, converter.ToStoryActivityResponse(created))
}

func (h *StoryActivityHandler) Update(c *gin.Context) {
	req := &dto.UpdateStoryActivityReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateStoryActivityUC.Execute(c.Request.Context(), converter.ToUpdateStoryActivityCommand(req))
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToStoryActivityResponse(updated))
}

func (h *StoryActivityHandler) Delete(c *gin.Context) {
	req := &dto.DeleteStoryActivityReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd := converter.ToDeleteStoryActivityCommand(req)

	if err := h.deleteStoryActivityUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
