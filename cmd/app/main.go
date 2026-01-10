package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lubosgarancovsky/eden-inri/docs"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres"
	"github.com/lubosgarancovsky/eden-inri/internal/app"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
)

var ServiceName = "inri-service"

// @securityDefinitions.apikey GatewayAuth
// @description Injected by API Gateway. Do not use in production.
// @in header
// @name X-User-ID
func main() {
	fmt.Println("███████╗██████╗ ███████╗███╗   ██╗      ██╗███╗   ██╗██████╗ ██╗\n██╔════╝██╔══██╗██╔════╝████╗  ██║      ██║████╗  ██║██╔══██╗██║\n█████╗  ██║  ██║█████╗  ██╔██╗ ██║█████╗██║██╔██╗ ██║██████╔╝██║\n██╔══╝  ██║  ██║██╔══╝  ██║╚██╗██║╚════╝██║██║╚██╗██║██╔══██╗██║\n███████╗██████╔╝███████╗██║ ╚████║      ██║██║ ╚████║██║  ██║██║\n╚══════╝╚═════╝ ╚══════╝╚═╝  ╚═══╝      ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝\n                                                                ")

	config.Init()

	dbconn, err := postgres.ConnectDB(config.GlobalConfig.DBUrl)
	if err != nil {
		log.Println("DB unavailable, starting without DB:", err)
	}

	container := app.NewContainer(dbconn)
	r := http.NewServerRoute(container)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err = r.Run(fmt.Sprintf(":%d", config.GlobalConfig.Port)); err != nil {
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
