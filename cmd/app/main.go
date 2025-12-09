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

func main() {
	r := gin.Default()
	cfg := config.LoadConfig()
	dbconn := db.ConnectDB(cfg.DBUrl)

	router.SetupRouter(r, cfg, dbconn)

	err := r.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
}
