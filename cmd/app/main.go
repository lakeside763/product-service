package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/product-service/config"
	"github.com/lakeside763/product-service/internal/adapters/cache"
	"github.com/lakeside763/product-service/internal/adapters/repositories"
	"github.com/lakeside763/product-service/internal/core/services"
	"github.com/lakeside763/product-service/internal/ports/http/handlers"
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

	// Initialize Data repositories
	dataRepo, err := repositories.NewDataRepo(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Error initializing repositories: %v", err)
	}

	cache := cache.NewRedisCache(config.RedisURL)

	productService := services.NewProductService(dataRepo.Product, cache) // Initialize product service
	productHandler := handlers.NewProductHandler(productService) // Initialize product handler

	// Main router
	router := httprouter.New()

	// Register Routes
	routes.ProductRouter(router, productHandler)

	log.Printf("Server is running on port %d", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}