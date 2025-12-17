package handler

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/listing"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

type ProjectHandler struct {
	s      *service.ProjectService
	parser *rsql.Parser
}

type ProjectPage struct {
	Items      []model.Project
	Page       int
	PageSize   int
	TotalCount int64
}

type MemberPage struct {
	Items      []model.ProjectUser
	Page       int
	PageSize   int
	TotalCount int64
}

func NewProjectHandler(parser *rsql.Parser, s *service.ProjectService) *ProjectHandler {
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
func (h *ProjectHandler) FindAll(c *gin.Context) {
	lq, apiErr := helpers.CreateListingQuery(c, h.parser, listing.ProjectFilter, listing.ProjectSort)
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

// FindByID @Summary      Get project by ID
// @Description  Returns a project by its ID
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {object}   model.Project
// @Router       /v1/inri/projects/{projectId} [get]
func (h *ProjectHandler) FindByID(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "projectId")
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

// Create @Summary      Create a new project
// @Description  Creates a new project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project  body  model.ProjectRequest  true  "Project data"
// @Success      201  {object}  model.Project
// @Router       /v1/inri/projects [post]
func (h *ProjectHandler) Create(c *gin.Context) {
	var input model.ProjectRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Create(user.ID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, result)
}

// Update @Summary      Update a project
// @Description  Updates an existing project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project  body  model.ProjectRequest  true  "Project data"
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {object}  model.Project
// @Router       /v1/inri/projects/{projectId} [put]
func (h *ProjectHandler) Update(c *gin.Context) {
	var input model.ProjectRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	UID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Update(user.ID, UID, &input)
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
func (h *ProjectHandler) Delete(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	_, err = h.s.Delete(user.ID, UID)
	if err != nil {
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
func (h *ProjectHandler) UploadAttachments(c *gin.Context) {
	projectID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

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

	if err := h.s.SaveAttachments(c, user.ID, projectID, files); err != nil {
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
// @Success      200  {object}   []model.Attachment
// @Router       /v1/inri/projects/{projectId}/attachments [get]
func (h *ProjectHandler) ListAttachments(c *gin.Context) {
	projectID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.ListAttachments(user.ID, projectID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// ListMembers @Summary      List project members
// @Description  Returns a paginated list of all project members
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}   MemberPage
// @Router       /v1/inri/projects/{projectId}/members [get]
func (h *ProjectHandler) ListMembers(c *gin.Context) {
	projectID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	lq, apiErr := helpers.CreateListingQuery(c, h.parser, listing.ProjectUserFilter, listing.ProjectUserSort)
	if apiErr != nil {
		c.Error(apiErr)
		return
	}

	page, err := h.s.ListProjectMembers(user.ID, projectID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, page)
}

// RemoveMember @Summary      Remove a project member
// @Description  Removes a member from a project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        memberId    path      string  true  "Member User ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/projects/{projectId}/members/{memberId} [delete]
func (h *ProjectHandler) RemoveMember(c *gin.Context) {
	projectID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}

	memberID, err := helpers.ExtractID(c, "memberId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.s.RemoveProjectMember(user.ID, memberID, projectID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
