package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/product-service/internal/core/models"
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
	cursorId := queryParams.Get("cursorId")
	pageSize := queryParams.Get("pageSize")

	priceLessThanToInt, _ := strconv.Atoi(priceLessThan)
	pageSizeToInt, _ := strconv.Atoi(pageSize)

	// Check if category is empty
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	products,nextCursorId, err := h.productService.GetProductsWithDiscount(category, priceLessThanToInt, cursorId, pageSizeToInt)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	responseData := models.ProductDataResponse{
		Data: products,
		CursorId: nextCursorId,
	}

	utils.JSONResponse(w, http.StatusOK, responseData)	
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Parse the request body
	var input models.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		// http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		utils.JSONErrorResponse(w, http.StatusBadRequest, err, "invalid input: %v")
		return
	}

	// validate the input
	if input.Name == "" || input.Category == "" || input.Price <= 0 {
		// http.Error(w, `{"error": "Invalid product details"}`, http.StatusBadRequest)
		utils.JSONErrorResponse(w, http.StatusBadRequest, nil, "invalid product details: %v")
		return
	}

	// Call the product service
	product, err := h.productService.CreateNewProduct(input)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err, "failed to create product: %v")
		return
	}

	utils.JSONResponse(w, http.StatusOK, product)
}

// w.Header().Set("Content-Type", "application/json")
// json.NewEncoder(w).Encode(products)