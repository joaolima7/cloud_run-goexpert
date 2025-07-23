package main

import (
	"log"

	"github.com/joaolima7/cloud_run-goexpert/config"
	"github.com/joaolima7/cloud_run-goexpert/internal/app"
	"github.com/joaolima7/cloud_run-goexpert/internal/infra/web/handler"
	"github.com/joaolima7/cloud_run-goexpert/internal/infra/web/webserver"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	getCityUC, err := app.InitializeGetCityByCepUseCase()
	if err != nil {
		log.Fatalf("Error initializing GetCityByCepUseCase: %v", err)
	}

	weatherUC, err := app.InitializeGetWeatherByCityUseCase()
	if err != nil {
		log.Fatalf("Error initializing GetWeatherByCityUseCase: %v", err)
	}

	server := webserver.NewWebServer(":" + cfg.WebServerPort)

	appHandler := handler.NewAppHandler(getCityUC, weatherUC)
	appHandler.RegisterRoutes(server.Router)

	log.Printf("Starting server on port %s", cfg.WebServerPort)
	server.Start()
}
