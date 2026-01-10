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

var ProjectListConfig = types.ListConfig{
	Filter: map[string]string{
		"name":           "name",
		"status":         "status",
		"slug":           "slug",
		"lastActivityAt": "last_activity_at",
		"isStarred":      "pu.is_starred",
		"role":           "pu.role",
	},
	Sort: map[string]string{
		"name":           "name",
		"createdAt":      "created_at",
		"updatedAt":      "updated_at",
		"lastActivityAt": "last_activity_at",
	},
}

type ProjectHandler struct {
	s      *services.ProjectService
	parser *rsql.Parser
}

type ProjectPage struct {
	Items      []models.Project
	Page       int
	PageSize   int
	TotalCount int64
}

func NewProjectHandler(parser *rsql.Parser, s *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{s, parser}
}

// FindAll @Summary      List projects
// @Description  Returns a paginated list of all projects
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}   ProjectPage
// @Router       /v1/inri/projects [get]
// @security GatewayAuth
func (h *ProjectHandler) FindAll(c *gin.Context) {
	helpers.HandleList(c, h.parser, ProjectListConfig, h.s.FindAll)
}

// FindByID @Summary      Get project by ID
// @Description  Returns a project by its ID
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {object}   models.Project
// @Router       /v1/inri/projects/{projectId} [get]
// @security GatewayAuth
func (h *ProjectHandler) FindByID(c *gin.Context) {
	helpers.HandleFindByID(c, "projectId", h.s.FindByID)
}

// Create @Summary      Create a new project
// @Description  Creates a new project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project  body  models.ProjectRequest  true  "Project data"
// @Success      201  {object}  models.Project
// @Router       /v1/inri/projects [post]
// @security GatewayAuth
func (h *ProjectHandler) Create(c *gin.Context) {
	helpers.HandleCreate(c, h.s.Create)
}

// Update @Summary      Update a project
// @Description  Updates an existing project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project  body  models.ProjectRequest  true  "Project data"
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {object}  models.Project
// @Router       /v1/inri/projects/{projectId} [put]
// @security GatewayAuth
func (h *ProjectHandler) Update(c *gin.Context) {
	var input models.ProjectRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	projectID := helpers.ExtractID(c, "projectId")

	result, err := h.s.Update(c.Request.Context(), projectID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// Delete @Summary      Delete a project
// @Description  Deletes the project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/projects/{projectId} [delete]
// @security GatewayAuth
func (h *ProjectHandler) Delete(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")

	if err := h.s.Delete(c.Request.Context(), projectID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

// UploadAttachments @Summary Upload attachments to a project
// @Description Upload multiple files as attachments for the given project
// @Tags         Projects
// @Accept       mpfd
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param files formData []file true "Files to upload"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/projects/{projectId}/attachments [post]
// @security GatewayAuth
func (h *ProjectHandler) UploadAttachments(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")

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

	if err := h.s.SaveAttachments(c, projectID, files); err != nil {
		c.Error(err)
		return
	}
	c.Status(204)
}

// ListAttachments @Summary      List project attachments
// @Description  Returns a list of attachments by project ID
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Success      200  {object}   []models.Attachment
// @Router       /v1/inri/projects/{projectId}/attachments [get]
// @security GatewayAuth
func (h *ProjectHandler) ListAttachments(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")

	result, err := h.s.ListAttachments(c.Request.Context(), projectID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// DownloadAttachment @Summary     Download attachment
// @Description  Downloads an attachment by ID
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Success      200  {object}   []models.Attachment
// @Router       /v1/inri/projects/{projectId}/attachments/{attachmentId} [get]
// @security GatewayAuth
func (h *ProjectHandler) DownloadAttachment(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	attachmentID := helpers.ExtractID(c, "attachmentId")

	if err := h.s.DownloadAttachment(c, projectID, attachmentID); err != nil {
		c.Error(err)
		return
	}
}

// DeleteAttachment @Summary      Delete attachment
// @Description  Deletes an attachment from the project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Success      204  {string}   "No content"
// @Router       /v1/inri/projects/{projectId}/attachments/{attachmentId} [delete]
// @security GatewayAuth
func (h *ProjectHandler) DeleteAttachment(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	attachmentID := helpers.ExtractID(c, "attachmentId")

	if err := h.s.DeleteAttachment(c.Request.Context(), projectID, attachmentID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

// Favourite @Summary      Add/Remove from favorites
// @Description  Toggles the is_starred flag in a project for current user
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Project
// @Router       /v1/inri/projects/:projectId/favourite [put]
// @security GatewayAuth
func (h *ProjectHandler) Favourite(c *gin.Context) {
	userID := helpers.GetUserContext(c).ID
	projectID := helpers.ExtractID(c, "projectId")

	result, err := h.s.Favourite(c.Request.Context(), userID, projectID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, result)
}
