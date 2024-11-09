package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/product-service/internal/ports/http/handlers"
)

func ProductRouter(router *httprouter.Router, productHandler *handlers.ProductHandler) {
	router.GET("/products", productHandler.GetProducts)
}


// ProductRouter initializes product-related routes and dependencies
/*
func ProductRouter(router *httprouter.Router, db *gorm.DB) {
	// Initialize product repository
	productRepo := repositories.NewProductRepo(db)

	// Initialize product service
	productService := services.NewProductService(productRepo)

	// Initialize product handler
	productHandler := handlers.NewProductHandler(productService)

	// Register routes
	router.GET("/products", productHandler.GetProducts)
}
	*/