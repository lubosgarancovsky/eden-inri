package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/middleware"
	"github.com/lubosgarancovsky/eden-inri/internal/app"
)

func NewServerRoute(c *app.Container) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.ErrorHandlerMiddleware())

	v1 := router.Group("/v1/inri")
	protected := v1.Group("", middleware.AuthMiddleware())

	{
		clients := protected.Group("/clients")
		{
			clients.POST("", c.ClientHandler.Create)
			clients.PUT("/:clientId", c.ClientHandler.Update)
			clients.DELETE("/:clientId", c.ClientHandler.Delete)
			clients.GET("/:clientId", c.ClientHandler.FindByID)
			clients.GET("", c.ClientHandler.List)
		}

		invoices := protected.Group("/invoices")
		{
			invoices.POST("", c.InvoiceHandler.Create)
			invoices.PUT("/:invoiceId", c.InvoiceHandler.Update)
			invoices.DELETE("/:invoiceId", c.InvoiceHandler.Delete)
			invoices.GET("/:invoiceId", c.InvoiceHandler.FindByID)
			invoices.GET("", c.InvoiceHandler.List)
		}
	}

	return router
}
