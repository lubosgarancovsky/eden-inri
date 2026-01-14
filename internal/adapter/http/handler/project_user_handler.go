package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/go-kit"
)

type ProjectUserHandler struct {
	listUC       ports.ListProjectUsersUseCase
	deleteUC     ports.DeleteProjectUserUseCase
	changeRoleUC ports.ChangeProjectUserRoleUseCase
	parser       *go_kit.Parser
}

func NewProjectUserHandler(
	listUC ports.ListProjectUsersUseCase,
	deleteUC ports.DeleteProjectUserUseCase,
	changeRoleUC ports.ChangeProjectUserRoleUseCase,
	parser *go_kit.Parser,
) *ProjectUserHandler {
	return &ProjectUserHandler{
		listUC,
		deleteUC,
		changeRoleUC,
		parser,
	}
}

type ProjectUserListingAttributes struct {
	IsStarred bool   `rsql:"filter"`
	Role      string `rsql:"filter,sort"`
	JoinedAt  string `rsql:"filter,sort"`
	FirstName string `rsql:"field:User.first_name,filter,sort"`
	LastName  string `rsql:"field:User.last_name,filter,sort"`
	Username  string `rsql:"field:User.username,filter,sort"`
}

func (h *ProjectUserHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &ProjectUserListingAttributes{})

	req := &dto.ListProjectUsersReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToScopedListQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	items, total, err := h.listUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(items, converter.ToProjectUserResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *ProjectUserHandler) Delete(c *gin.Context) {
	req := &dto.DeleteProjectUserReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToScopedCommand(req)
	if err != nil {
		handle.Error(c, err)
	}

	if err = h.deleteUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ProjectUserHandler) ChangeRole(c *gin.Context) {
	req := &dto.ChangeProjectUserRoleReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToChangeProjectUserRoleCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	projectUser, err := h.changeRoleUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToProjectUserResponse(projectUser))
}
