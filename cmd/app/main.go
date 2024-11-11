package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	server := setupServer(config.Port, router)

	startServer(server)
	waitForShutdown(server, dataRepo, cache)
}

func setupServer(port int, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
}

func startServer(server *http.Server) {
	go func() {
		log.Printf("Server is running on port %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
}

func closeResources(dataRepo *repositories.DataRepo, cache *cache.RedisCache) {
	if err := dataRepo.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
	cache.Close()
}

func waitForShutdown(server *http.Server, dataRepo *repositories.DataRepo, cache *cache.RedisCache) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	// Close database and cache resources
	closeResources(dataRepo, cache)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}