package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/product-service/config"
	"github.com/lakeside763/product-service/internal/adapters/cache"
	"github.com/lakeside763/product-service/internal/adapters/repositories"
	"github.com/lakeside763/product-service/internal/ports/http/routes"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		log.Info("Loading Configuration")
	}
}

func main() {
	// Initialize Configuration
	config := config.LoadConfig()

	// Initialize Data repositories and database
	dataRepo, err := repositories.NewDataRepo(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Error initializing repositories: %v", err)
	}

	cache := cache.NewRedisCache(config.RedisURL) // Initialize redis cache

	// Main router
	router := httprouter.New()

	// Register Routes
	// Register product routes for managing product handler and service
	routes.ProductRouter(router, dataRepo.Product, cache)

	log.Printf("Server is running on port %d", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}