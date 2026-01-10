package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

type ClientHandler struct {
	createClientUC ports.CreateClientUseCase
	updateClientUC ports.UpdateClientUseCase
	deleteClientUC ports.DeleteClientUseCase
}

func NewClientHandler(
	createClientUC ports.CreateClientUseCase,
	updateClientUC ports.UpdateClientUseCase,
	deleteClientUC ports.DeleteClientUseCase,
) *ClientHandler {
	return &ClientHandler{
		createClientUC,
		updateClientUC,
		deleteClientUC,
	}
}

func (h *ClientHandler) Create(c *gin.Context) {
	body := &dto.CreateClientReq{}

	if err := validator.BindAndValidate(c, body); err != nil {
		// TODO: Return actual error
		c.JSON(404, gin.H{
			"status": "handler fn ERROR",
		})
		return
	}

	created, err := h.createClientUC.Execute(c.Request.Context(), body.ToCommand())
	if err != nil {
		// TODO: Return actual error
		c.JSON(4500, gin.H{
			"status": "handler fn ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.ToClientResponse(created))
}
