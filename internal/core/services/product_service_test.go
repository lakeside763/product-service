package services

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/lakeside763/product-service/internal/core/models"
	"github.com/lakeside763/product-service/internal/ports/interfaces/mocks"
	"github.com/lakeside763/product-service/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetProductWithDiscount(t *testing.T) {
	// Create a new Gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock dependencies
	mockRepo := interfaces.NewMockProducts(ctrl)
	mockCache := interfaces.NewMockRedis(ctrl)

	// Create a prodcutService instance with mocks
	productService := NewProductService(mockRepo, mockCache)

	// Mock product data
	product := &models.Product{
		Sku:       "12345",
		Name:      "Test Product",
		Category:  "Electronics",
		Price:     10000, // 100.00 EUR in integer format
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Expected result
	expectedResponse := &models.ProductWithDiscountResponse{
		Sku:      product.Sku,
		Name:     product.Name,
		Category: product.Category,
		Price: models.PriceWithDiscount{
			Original:           utils.ConvertPriceToDisplayFormat(product.Price),
			Final:              utils.ConvertPriceToDisplayFormat(product.Price), // No discount applied
			DiscountPercentage: "",
			Currency:           "EUR",
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	// Set up mock expectations
	mockRepo.EXPECT().GetProducts(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Product{product}, nil)
	mockCache.EXPECT().Get(gomock.Any()).Return("", errors.New("cache miss"))
	mockRepo.EXPECT().GetMaxDiscount(product.Category, product.Sku).Return(0.0, nil) // No discount
	mockCache.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	// Execute the method
	result, err := productService.GetProductsWithDiscount(15000, "", 10)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, expectedResponse, result[0])
}

func TestGetProductWithDiscount_WithDiscountFromCache(t *testing.T) {
	// Create a new Gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock dependencies
	mockRepo := interfaces.NewMockProducts(ctrl)
	mockCache := interfaces.NewMockRedis(ctrl)

	// Create a ProductService instance with mocks
	productService := NewProductService(mockRepo, mockCache)

	// Mock product data
	product := &models.Product{
		Sku:       "12345",
		Name:      "Discounted Product",
		Category:  "Apparel",
		Price:     20000, // 200.00 EUR in integer format
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	discount := 20.0 // 20% discount
	finalPrice := int(float64(product.Price) * (1 - discount/100))

	// Expected result with discount
	expectedResponse := &models.ProductWithDiscountResponse{
		Sku:      product.Sku,
		Name:     product.Name,
		Category: product.Category,
		Price: models.PriceWithDiscount{
			Original:           utils.ConvertPriceToDisplayFormat(product.Price),
			Final:              utils.ConvertPriceToDisplayFormat(finalPrice),
			DiscountPercentage: "20%",
			Currency:           "EUR",
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	// Set up mock expectations
	mockRepo.EXPECT().GetProducts(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Product{product}, nil)
	mockCache.EXPECT().Get(gomock.Any()).Return("20.0", nil) // Cache hit with 20% discount

	// Execute the method
	result, err := productService.GetProductsWithDiscount(30000, "", 10)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, expectedResponse, result[0])
}