package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/app"
)

func NewServerRoute(c *app.Container) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	v1 := router.Group("/v1/inri")
	{
		clients := v1.Group("/clients")
		{
			clients.POST("", c.ClientHandler.Create)
		}
	}

	return router
}
