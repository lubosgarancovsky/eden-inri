package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/pkg/api"
)

type Server struct {
	router        *gin.Engine
	clientService api.ClientService
}

func NewServer(
	router *gin.Engine,
	clientService api.ClientService,
) *Server {
	return &Server{
		router:        router,
		clientService: clientService,
	}
}
