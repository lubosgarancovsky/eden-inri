package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/eden-inri/pkg/types"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

var ProjectUserListConfig = &types.ListConfig{
	Filter: map[string]string{
		"projectId":     "project_id",
		"userId":        "user_id",
		"role":          "role",
		"joinedAt":      "joined_at",
		"userFirstName": "User.first_name",
		"userLastName":  "User.first_name",
	},
	Sort: map[string]string{
		"role":          "role",
		"joinedAt":      "joined_at",
		"userFirstName": "User.first_name",
		"userLastName":  "User.first_name",
	},
}

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
	projectID := helpers.ExtractID(c, "projectId")
	lq := helpers.CreateListingQuery(c, h.parser, ProjectUserListConfig.Filter, ProjectUserListConfig.Sort)

	page, err := h.s.FindAll(c.Request.Context(), projectID, lq)
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
	projectID := helpers.ExtractID(c, "projectId")
	memberID := helpers.ExtractID(c, "memberId")
	userID := helpers.GetUserContext(c).ID

	var input struct {
		Role model.ProjectRole `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	updatedMember, err := h.s.Update(c.Request.Context(), userID, projectID, memberID, input.Role)
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
	projectID := helpers.ExtractID(c, "projectId")
	memberID := helpers.ExtractID(c, "memberId")

	if err := h.s.Delete(c.Request.Context(), memberID, projectID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
