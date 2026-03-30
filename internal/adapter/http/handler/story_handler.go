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
	"github.com/lubosgarancovsky/go-kit"
)

type StoryHandler struct {
	createStoryUC         ports.CreateStoryUseCase
	updateStoryUC         ports.UpdateStoryUseCase
	deleteStoryUC         ports.DeleteStoryUseCase
	findStoryByIDUC       ports.FindStoryByIDUseCase
	findStoryBySlugUC     ports.FindStoryBySlugUseCase
	listStoriesUC         ports.ListStoriesUseCase
	listAssignedStoriesUC ports.ListAssignedStoriesUseCase
	changeAssigneeUC      ports.ChangeStoryAssigneeUseCase
	parser                *go_kit.Parser
}

func NewStoryHandler(
	createStoryUC ports.CreateStoryUseCase,
	updateStoryUC ports.UpdateStoryUseCase,
	deleteStoryUC ports.DeleteStoryUseCase,
	findStoryByIDUC ports.FindStoryByIDUseCase,
	findStoryBySlugUC ports.FindStoryBySlugUseCase,
	listStoriesUC ports.ListStoriesUseCase,
	listAssignedStoriesUC ports.ListAssignedStoriesUseCase,
	changeAssigneeUC ports.ChangeStoryAssigneeUseCase,
	parser *go_kit.Parser,
) *StoryHandler {
	return &StoryHandler{
		createStoryUC,
		updateStoryUC,
		deleteStoryUC,
		findStoryByIDUC,
		findStoryBySlugUC,
		listStoriesUC,
		listAssignedStoriesUC,
		changeAssigneeUC,
		parser,
	}
}

type StoryListingAttributes struct {
	ColumnID       string `rsql:"filter"`
	Title          string `rsql:"filter,sort"`
	Kind           string `rsql:"filter"`
	CreatedAt      string `rsql:"filter,sort"`
	Priority       string `rsql:"filter,sort"`
	AssigneeID     string `rsql:"filter"`
	LastActivityAt string `rsql:"sort"`
}

func (h *StoryHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &StoryListingAttributes{})

	req := &dto.ListStoriesReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListStoriesQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	stories, total, err := h.listStoriesUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(stories, converter.ToStoryResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *StoryHandler) ListAssigned(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &StoryListingAttributes{})

	req := &dto.ListAssignedStoriesReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListAssignedStoriesQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	stories, total, err := h.listAssignedStoriesUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(stories, converter.ToStoryResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *StoryHandler) FindByID(c *gin.Context) {
	req := &dto.FindStoryByIDReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToFindStoryByIDQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	story, err := h.findStoryByIDUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToStoryResponse(story))
}

func (h *StoryHandler) FindBySlug(c *gin.Context) {
	req := &dto.FindStoryBySlugReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToFindStoryBySlugQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	story, err := h.findStoryBySlugUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
	}

	c.JSON(http.StatusOK, converter.ToStoryResponse(story))
}

func (h *StoryHandler) Create(c *gin.Context) {
	req := &dto.CreateStoryReq{}
	fmt.Println(c.Request.Header)

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToCreateStoryCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createStoryUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusCreated, converter.ToStoryResponse(created))
}

func (h *StoryHandler) Update(c *gin.Context) {
	req := &dto.UpdateStoryReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToUpdateStoryCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateStoryUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToStoryResponse(updated))
}

func (h *StoryHandler) Delete(c *gin.Context) {
	req := &dto.DeleteStoryReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToDeleteStoryCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.deleteStoryUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *StoryHandler) ChangeAssignee(c *gin.Context) {
	req := &dto.ChangeStoryAssigneeReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToChangeStoryAssigneeCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.changeAssigneeUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
