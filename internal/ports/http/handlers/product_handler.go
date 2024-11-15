package handlers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/product-service/internal/core/services"
	"github.com/lakeside763/product-service/pkg/utils"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: service}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryParams := r.URL.Query()
	category := queryParams.Get("category")
	priceLessThan := queryParams.Get("priceLessThan")
	lastProductId := queryParams.Get("lastProductId")
	pageSize := queryParams.Get("pageSize")

	priceLessThanToInt, _ := strconv.Atoi(priceLessThan)
	pageSizeToInt, _ := strconv.Atoi(pageSize)

	// Check if category is empty
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	products, err := h.productService.GetProductsWithDiscount(category, priceLessThanToInt, lastProductId, pageSizeToInt)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, http.StatusOK, products)	
}

// w.Header().Set("Content-Type", "application/json")
// json.NewEncoder(w).Encode(products)