package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/listing"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

type ProjectUserHandler struct {
	s      *service.ProjectUserService
	parser *rsql.Parser
}

type MemberPage struct {
	Items      []model.ProjectUser
	Page       int
	PageSize   int
	TotalCount int64
}

func NewProjectUserHandler(parser *rsql.Parser, s *service.ProjectUserService) *ProjectUserHandler {
	return &ProjectUserHandler{s, parser}
}

// FindAll @Summary      List project members
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
func (h *ProjectUserHandler) FindAll(c *gin.Context) {
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

	page, err := h.s.FindAll(user.ID, projectID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, page)
}

// Update @Summary      Update member role
// @Description  Updates the role of a project member
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        memberId    path      string  true  "Member User ID"
// @Param        body        body      model.UpdateProjectUserRequest true "New role"
// @Success      200  {object}  model.ProjectUser
// @Router       /v1/inri/projects/{projectId}/members/{memberId} [put]
func (h *ProjectUserHandler) Update(c *gin.Context) {
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

	var input struct {
		Role model.ProjectRole `json:"role" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	updatedMember, err := h.s.Update(user.ID, projectID, memberID, input.Role)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, updatedMember)
}

// Delete @Summary      Remove a project member
// @Description  Removes a member from a project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        memberId    path      string  true  "Member User ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/projects/{projectId}/members/{memberId} [delete]
func (h *ProjectUserHandler) Delete(c *gin.Context) {
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

	if err := h.s.Delete(user.ID, memberID, projectID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
