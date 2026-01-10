package handlers

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/services"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/eden-inri/pkg/types"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

var StoryListConfig = types.ListConfig{
	Filter: map[string]string{
		"id":         "id",
		"projectId":  "project_id",
		"columnId":   "column_id",
		"boardId":    "board_id",
		"slug":       "slug",
		"title":      "title",
		"kind":       "kind",
		"assigneeId": "assignee_id",
		"priority":   "priority",
		"startDate":  "start_date",
		"endDate":    "end_date",
		"createdAt":  "created_at",
	},
	Sort: map[string]string{
		"position":  "position",
		"priority":  "priority",
		"createdAt": "created_at",
		"updatedAt": "updated_at",
		"title":     "title",
		"slug":      "slug",
	},
}

type StoryHandler struct {
	parser *rsql.Parser
	s      *services.StoryService
}

type StoryPage struct {
	Items      []models.StoryListItem
	Page       int
	PageSize   int
	TotalCount int64
}

func NewStoryHandler(p *rsql.Parser, s *services.StoryService) *StoryHandler {
	return &StoryHandler{s: s, parser: p}
}

// FindAllAssigned @Summary      List stories assigned to current user
// @Description  Returns all stories that are assigned to the caller
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {array}  StoryPage
// @Router       /v1/inri/stories [get]
// @security GatewayAuth
func (h *StoryHandler) FindAllAssigned(c *gin.Context) {
	lq := helpers.CreateListingQuery(c, h.parser, StoryListConfig.Filter, StoryListConfig.Sort)
	userID := helpers.GetUserContext(c).ID

	stories, err := h.s.FindAllAssigned(c.Request.Context(), userID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, stories)
}

// FindAll @Summary      List stories in a column
// @Description  Returns all stories for a column
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        columnId   path      string  true  "Column ID"
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {array}  StoryPage
// @Router       /v1/inri/projects/{projectId}/stories [get]
// @security GatewayAuth
func (h *StoryHandler) FindAll(c *gin.Context) {
	lq := helpers.CreateListingQuery(c, h.parser, StoryListConfig.Filter, StoryListConfig.Sort)
	projectID := helpers.ExtractID(c, "projectId")

	stories, err := h.s.FindAll(c.Request.Context(), projectID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, stories)
}

// FindByID @Summary      Get story by ID
// @Description  Returns a specific story by ID
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        storyId    path      string  true  "Story ID"
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {object}  models.Story
// @Router       /v1/inri/projects/{projectId}/stories/{storyId} [get]
// @security GatewayAuth
func (h *StoryHandler) FindByID(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

	story, err := h.s.FindByID(c.Request.Context(), projectID, storyID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, story)
}

// FindBySlug @Summary      Get story by slug
// @Description  Returns a specific story by slug within a project
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        slug        path      string  true  "Story slug"
// @Success      200  {object}  models.Story
// @Router       /v1/inri/projects/{projectId}/stories/slug/{slug} [get]
// @security GatewayAuth
func (h *StoryHandler) FindBySlug(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	slug := c.Param("slug")

	story, err := h.s.FindBySlug(c.Request.Context(), projectID, slug)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, story)
}

// Insert @Summary      Create a new story
// @Description  Creates a new story in a column
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        body       body      models.StoryRequest true "Story data"
// @Success      201  {object}  models.Story
// @Router       /v1/inri/projects/{projectId}/stories [post]
// @security GatewayAuth
func (h *StoryHandler) Insert(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	userID := helpers.GetUserContext(c).ID

	var req models.StoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	story, err := h.s.Insert(c.Request.Context(), userID, projectID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, story)
}

// Update @Summary      Update a story
// @Description  Updates an existing story
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        storyId    path      string  true  "Story ID"
// @Param        projectId   path      string  true  "Project ID"
// @Param        body       body      models.StoryRequest true "Updated story data"
// @Success      200  {object}  models.Story
// @Router       /v1/inri/projects/{projectId}/stories/{storyId} [put]
// @security GatewayAuth
func (h *StoryHandler) Update(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

	var req models.StoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	story, err := h.s.Update(c.Request.Context(), projectID, storyID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, story)
}

// Delete @Summary      Delete a story
// @Description  Deletes a story
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        storyId    path      string  true  "Story ID"
// @Param        projectId   path      string  true  "Project ID"
// @Success      204  {string} string "No Content"
// @Router       /v1/inri/projects/{projectId}/stories/{storyId} [delete]
// @security GatewayAuth
func (h *StoryHandler) Delete(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

	if err := h.s.Delete(c.Request.Context(), projectID, storyID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

// GetAssignee @Summary      Get story assignee
// @Description  Returns an assignee of a story
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        storyId    path      string  true  "Story ID"
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {object} models.User
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/assignee [get]
// @security GatewayAuth
func (h *StoryHandler) GetAssignee(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

	user, err := h.s.GetAssignee(c.Request.Context(), projectID, storyID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, user)
}

// ChangeAssignee @Summary      Change an assignee
// @Description  Changes an assignee of a story
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        storyId    path      string  true  "Story ID"
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {object} models.StoryAssigneeRequest
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/assignee [put]
// @security GatewayAuth
func (h *StoryHandler) ChangeAssignee(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

	var input models.StoryAssigneeRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	result, err := h.s.ChangeAssignee(c.Request.Context(), projectID, storyID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// UploadAttachments @Summary Upload attachments to a story
// @Description Upload multiple files as attachments for the given story
// @Tags         Kanban Stories
// @Accept       mpfd
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        storyId   path      string  true  "Story ID"
// @Param files formData []file true "Files to upload"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/attachments [post]
// @security GatewayAuth
func (h *StoryHandler) UploadAttachments(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

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

	if err := h.s.SaveAttachments(c, projectID, storyID, files); err != nil {
		c.Error(err)
		return
	}
	c.Status(204)
}

// ListAttachments @Summary      List story attachments
// @Description  Returns a list of attachments by story ID
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Success      200  {object}   []models.Attachment
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/attachments [get]
// @security GatewayAuth
func (h *StoryHandler) ListAttachments(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

	result, err := h.s.ListAttachments(c.Request.Context(), projectID, storyID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// DownloadAttachment @Summary     Download attachment
// @Description  Downloads an attachment by ID
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Success      200  {object}   []models.Attachment
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/attachments/{attachmentId} [get]
// @security GatewayAuth
func (h *StoryHandler) DownloadAttachment(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	attachmentID := helpers.ExtractID(c, "attachmentId")

	if err := h.s.DownloadAttachment(c, projectID, attachmentID); err != nil {
		c.Error(err)
		return
	}
}

// DeleteAttachment @Summary      Delete attachment
// @Description  Deletes an attachment from the project
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Success      204  {string}   "No content"
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/attachments/{attachmentId} [delete]
// @security GatewayAuth
func (h *StoryHandler) DeleteAttachment(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	attachmentID := helpers.ExtractID(c, "attachmentId")

	if err := h.s.DeleteAttachment(c.Request.Context(), projectID, attachmentID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
