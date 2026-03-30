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

type ClientHandler struct {
	createClientUC   ports.CreateClientUseCase
	updateClientUC   ports.UpdateClientUseCase
	deleteClientUC   ports.DeleteClientUseCase
	findClientByIDUC ports.FindClientByIDUseCase
	listClientsUC    ports.ListClientsUseCase
	parser           *go_kit.Parser
}

func NewClientHandler(
	createClientUC ports.CreateClientUseCase,
	updateClientUC ports.UpdateClientUseCase,
	deleteClientUC ports.DeleteClientUseCase,
	findClientByIDUC ports.FindClientByIDUseCase,
	listClientsUC ports.ListClientsUseCase,
	parser *go_kit.Parser,
) *ClientHandler {
	return &ClientHandler{
		createClientUC,
		updateClientUC,
		deleteClientUC,
		findClientByIDUC,
		listClientsUC,
		parser,
	}
}

type ClientListingAttributes struct {
	Name         string `rsql:"filter,sort"`
	ClientType   string `rsql:"filter"`
	ContractType string `rsql:"filter"`
	CreatedAt    string `rsql:"filter,sort"`
	StartedAt    string `rsql:"filter,sort"`
	FinishedAt   string `rsql:"filter,sort"`
}

func (h *ClientHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &ClientListingAttributes{})

	req := &dto.ListClientsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	clients, total, err := h.listClientsUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(clients, converter.ToClientResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *ClientHandler) FindByID(c *gin.Context) {
	req := &dto.FindClientByIDReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	client, err := h.findClientByIDUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToClientResponse(client))
}

func (h *ClientHandler) Create(c *gin.Context) {
	req := &dto.CreateClientReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToCreateClientCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createClientUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusCreated, converter.ToClientResponse(created))
}

func (h *ClientHandler) Update(c *gin.Context) {
	req := &dto.UpdateClientReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToUpdateClientCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateClientUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToClientResponse(updated))
}

func (h *ClientHandler) Delete(c *gin.Context) {
	req := &dto.DeleteClientReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.deleteClientUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
