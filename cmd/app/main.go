package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	fmt.Println("▗▄▄▄▖▗▄▄▄ ▗▄▄▄▖▗▖  ▗▖        ▗▄▄▄▖▗▖  ▗▖▗▄▄▖ ▗▄▄▄▖\n▐▌   ▐▌  █▐▌   ▐▛▚▖▐▌          █  ▐▛▚▖▐▌▐▌ ▐▌  █  \n▐▛▀▀▘▐▌  █▐▛▀▀▘▐▌ ▝▜▌          █  ▐▌ ▝▜▌▐▛▀▚▖  █  \n▐▙▄▄▖▐▙▄▄▀▐▙▄▄▖▐▌  ▐▌        ▗▄█▄▖▐▌  ▐▌▐▌ ▐▌▗▄█▄▖\n                                                  \n                                                  \n                                                  \n")

	r := gin.Default()
	cfg := config.LoadConfig()
	dbconn, err := db.ConnectDB(cfg.DBUrl)
	if err != nil {
		log.Println("DB unavailable, starting without DB:", err)
	}

	router.SetupRouter(r, cfg, dbconn)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := r.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
			log.Println(err)
			stop()
		}
	}()

	<-ctx.Done() // wait for shutdown signal

	log.Println("shutting down")

	if dbconn != nil {
		sqlDB, _ := dbconn.DB()
		sqlDB.Close()
	}
}
