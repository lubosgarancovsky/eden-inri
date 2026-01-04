package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lubosgarancovsky/eden-inri/docs"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/db"
	"github.com/lubosgarancovsky/eden-inri/internal/router"
)

// @securityDefinitions.apikey GatewayAuth
// @description Injected by API Gateway. Do not use in production.
// @in header
// @name X-User-ID
func main() {
	r := gin.Default()
	cfg := config.LoadConfig()
	dbconn, err := db.ConnectDB(cfg.DBUrl)
	if err != nil {
		log.Println("DB unavailable, starting without DB:", err)
	}

	router.SetupRouter(r, cfg, dbconn)

	err = r.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
}
