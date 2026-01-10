package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/services"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/eden-inri/pkg/types"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

var ProjectDocumentListConfig = &types.ListConfig{
	Filter: map[string]string{
		"name": "name",
	},
	Sort: map[string]string{
		"createdAt": "created_at",
		"updatedAt": "updated_at",
		"name":      "name",
	},
}

type ProjectDocumentHandler struct {
	s      *services.ProjectDocumentService
	parser *rsql.Parser
}

type ProjectDocumentPage struct {
	Items      []models.ProjectDocument
	Page       int
	PageSize   int
	TotalCount int64
}

func NewProjectDocumentHandler(parser *rsql.Parser, s *services.ProjectDocumentService) *ProjectDocumentHandler {
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
// @security GatewayAuth
func (h *ProjectDocumentHandler) FindAll(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	lq := helpers.CreateListingQuery(c, h.parser, ProjectDocumentListConfig.Filter, ProjectDocumentListConfig.Sort)

	result, err := h.s.FindAll(c.Request.Context(), projectID, lq)
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
// @Success      200  {object}   models.ProjectDocument
// @Router       /v1/inri/projects/{projectId}/documents/{documentId} [get]
// @security GatewayAuth
func (h *ProjectDocumentHandler) FindByID(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	documentID := helpers.ExtractID(c, "documentId")

	result, err := h.s.FindByID(c.Request.Context(), projectID, documentID)
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
// @Param        document  body  models.ProjectDocumentRequest  true  "Document data"
// @Success      201  {object}  models.ProjectDocument
// @Router       /v1/inri/projects/{projectId}/documents [post]
// @security GatewayAuth
func (h *ProjectDocumentHandler) Create(c *gin.Context) {
	var input models.ProjectDocumentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	projectID := helpers.ExtractID(c, "projectId")

	result, err := h.s.Create(c.Request.Context(), projectID, &input)
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
// @Param        document  body  models.ProjectDocumentRequest  true  "Document data"
// @Success      200  {object}  models.ProjectDocument
// @Router       /v1/inri/projects/{projectId}/documents/{documentId} [put]
// @security GatewayAuth
func (h *ProjectDocumentHandler) Update(c *gin.Context) {
	var input models.ProjectDocumentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	projectID := helpers.ExtractID(c, "projectId")
	documentID := helpers.ExtractID(c, "documentId")

	result, err := h.s.Update(c.Request.Context(), projectID, documentID, &input)
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
// @security GatewayAuth
func (h *ProjectDocumentHandler) Delete(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	documentID := helpers.ExtractID(c, "documentId")

	if err := h.s.Delete(c.Request.Context(), projectID, documentID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
