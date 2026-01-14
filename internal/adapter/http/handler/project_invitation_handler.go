package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

type ProjectInvitationHandler struct {
	inviteUC ports.InviteProjectUserPort
	acceptUC ports.AcceptProjectInvitationPort
}

func NewProjectInvitationHandler(inviteUC ports.InviteProjectUserPort, acceptUC ports.AcceptProjectInvitationPort) *ProjectInvitationHandler {
	return &ProjectInvitationHandler{inviteUC: inviteUC, acceptUC: acceptUC}
}

func (h *ProjectInvitationHandler) Invite(c *gin.Context) {
	req := &dto.ProjectInvitationReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	command, err := converter.ToProjectInvitationCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.inviteUC.Execute(c.Request.Context(), command); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ProjectInvitationHandler) Accept(c *gin.Context) {
	req := &dto.AcceptInvitationReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	command, err := converter.ToAcceptInvitationCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	invitation, err := h.acceptUC.Execute(c.Request.Context(), command)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToProjectInvitationResponse(invitation))
}
