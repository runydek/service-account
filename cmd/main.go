package main

import (
	"log"
	"os"
	"service-account/config"
	"service-account/migrations"
	"service-account/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file, ensure environment variables are set.")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := migrations.Migrate(cfg.DB); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	e := echo.New()

	routes.InitRoutes(e, cfg.DB)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	e.Logger.Fatal(e.Start(":" + port))
}
