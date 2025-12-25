package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
)

type ProjectInvitationsHandler struct {
	service *service.ProjectInvitationService
}

func NewProjectInvitationsHandler(service *service.ProjectInvitationService) *ProjectInvitationsHandler {
	return &ProjectInvitationsHandler{service}
}

// Create @Summary      Create an invitation
// @Description  Invites user to collaborate on a project
// @Tags         Project invitations
// @Accept       json
// @Produce      json
// @Param        invitation  body  model.ProjectInvitationRequest  true  "Project invitation data"
// @Success      204  {string}  "No content"
// @Router       /v1/inri/projects/{projectId}/invite [post]
// @security GatewayAuth
func (h *ProjectInvitationsHandler) Create(c *gin.Context) {
	var req model.ProjectInvitationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	userID := helpers.GetUserContext(c).ID
	projectID := helpers.ExtractID(c, "projectId")

	if err := h.service.Create(c.Request.Context(), userID, projectID, &req); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, gin.H{})
}

// Accept @Summary      Accpet an invitation
// @Description  Accepts the invitation to collaborate on a project
// @Tags         Project invitations
// @Accept       json
// @Produce      json
// @Param        invitation  body  model.AcceptInvitationRequest  true  "Accept request body"
// @Success      200  {object} model.ProjectInvitation
// @Router       /v1/inri/projects/accept-invitation [post]
// @security GatewayAuth
func (h *ProjectInvitationsHandler) Accept(c *gin.Context) {
	var req model.AcceptInvitationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	userID := helpers.GetUserContext(c).ID

	invitation, err := h.service.Accept(c.Request.Context(), userID, req.Token)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, invitation)
}
