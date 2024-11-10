package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/product-service/internal/core/services"
	"github.com/lakeside763/product-service/internal/ports/http/handlers"
	"github.com/lakeside763/product-service/internal/ports/interfaces"
)

// ProductRouter initializes product-related routes and dependencies
func ProductRouter(router *httprouter.Router, productRepo interfaces.Products, cache interfaces.Redis) {

	// Initialize product service
	productService := services.NewProductService(productRepo, cache)
	// Initialize product handler
	productHandler := handlers.NewProductHandler(productService)

	// Register routes
	router.GET("/products", productHandler.GetProducts)
}
