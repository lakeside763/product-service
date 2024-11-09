package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/product-service/internal/core/services"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: service}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}