package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/listing"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

type ProjectDocumentHandler struct {
	s      *service.ProjectDocumentService
	parser *rsql.Parser
}

type ProjectDocumentPage struct {
	Items      []model.ProjectDocument
	Page       int
	PageSize   int
	TotalCount int64
}

func NewProjectDocumentHandler(parser *rsql.Parser, s *service.ProjectDocumentService) *ProjectDocumentHandler {
	return &ProjectDocumentHandler{s, parser}
}

// FindAll @Summary      List project documents
// @Description  Returns a paginated list of project documents for a project
// @Tags         ProjectDocuments
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}   ProjectDocumentPage
// @Router       /v1/inri/projects/{projectId}/documents [get]
func (h *ProjectDocumentHandler) FindAll(c *gin.Context) {
	lq, apiErr := helpers.CreateListingQuery(c, h.parser, listing.ProjectDocumentFilter, listing.ProjectDocumentSort)
	if apiErr != nil {
		c.Error(apiErr)
		return
	}

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

	result, err := h.s.FindAll(user.ID, projectID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// FindByID @Summary      Get project document by ID
// @Description  Returns a project document by its ID within the project
// @Tags         ProjectDocuments
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        documentId   path      string  true  "Document ID"
// @Success      200  {object}   model.ProjectDocument
// @Router       /v1/inri/projects/{projectId}/documents/{documentId} [get]
func (h *ProjectDocumentHandler) FindByID(c *gin.Context) {
	projectID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}
	UID, err := helpers.ExtractID(c, "documentId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.FindByID(user.ID, projectID, UID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// Create @Summary      Create a new project document
// @Description  Creates a new project document under the project
// @Tags         ProjectDocuments
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        document  body  model.ProjectDocumentRequest  true  "Document data"
// @Success      201  {object}  model.ProjectDocument
// @Router       /v1/inri/projects/{projectId}/documents [post]
func (h *ProjectDocumentHandler) Create(c *gin.Context) {
	var input model.ProjectDocumentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

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

	result, err := h.s.Create(user.ID, projectID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, result)
}

// Update @Summary      Update a project document
// @Description  Updates an existing project document
// @Tags         ProjectDocuments
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        documentId   path      string  true  "Document ID"
// @Param        document  body  model.ProjectDocumentRequest  true  "Document data"
// @Success      200  {object}  model.ProjectDocument
// @Router       /v1/inri/projects/{projectId}/documents/{documentId} [put]
func (h *ProjectDocumentHandler) Update(c *gin.Context) {
	var input model.ProjectDocumentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	projectID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}
	UID, err := helpers.ExtractID(c, "documentId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Update(user.ID, projectID, UID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// Delete @Summary      Delete a project document
// @Description  Deletes the project document
// @Tags         ProjectDocuments
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        documentId   path      string  true  "Document ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/projects/{projectId}/documents/{documentId} [delete]
func (h *ProjectDocumentHandler) Delete(c *gin.Context) {
	projectID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}
	UID, err := helpers.ExtractID(c, "documentId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	_, err = h.s.Delete(user.ID, projectID, UID)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
