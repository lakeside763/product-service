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
	// Using setup function to inialize dependencies
	setup := setupProductServiceTest(t)
	defer setup.Controller.Finish()

	product := createMockProduct("60d4403f-7ab7-4336-9017-0b397f71065f", "000001", "Leather ankle boots", "boots", 10000)
	expectedResponse := createExpectedResponse(product, 10000, "")

	// Set up mock expectations
	setup.MockRepo.EXPECT().GetProducts(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Product{product}, nil)
	setup.MockCache.EXPECT().Get(gomock.Any()).Return("", errors.New("cache miss"))
	setup.MockRepo.EXPECT().GetMaxDiscount(product.Category, product.Sku).Return(0.0, nil) // No discount
	setup.MockCache.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	// Execute the method
	result, err := setup.ProductService.GetProductsWithDiscount("boots", 15000, "", 10)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, expectedResponse, result[0])
}

func TestGetProductWithDiscount_WithDiscountFromCache(t *testing.T) {
	// Setup product service test
	setup := setupProductServiceTest(t)
	defer setup.Controller.Finish()

	// Now access mocks and services from the setup struct
	product := createMockProduct("60d4403f-7ab7-4336-9017-0b397f71065f", "12345", "Test Product", "Electronics", 10000)
	discount := 20.0 // 20% discount
	finalPrice := int(float64(product.Price) * (1 - discount/100))
	expectedResponse := createExpectedResponse(product, finalPrice, "20%")
	
	// Set up mock expectations
	setup.MockRepo.EXPECT().GetProducts(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Product{product}, nil)
	setup.MockCache.EXPECT().Get(gomock.Any()).Return("20.0", nil) // Cache hit with 20% discount

	// Execute the method
	result, err := setup.ProductService.GetProductsWithDiscount("boots", 30000, "", 10)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, expectedResponse, result[0])
	assert.Equal(t, expectedResponse.Name, result[0].Name)
}

type ProductServiceTest struct {
	Controller *gomock.Controller
	MockRepo *interfaces.MockProducts
	MockCache *interfaces.MockRedis
	ProductService *ProductService
}

func setupProductServiceTest(t *testing.T) *ProductServiceTest {
	ctrl := gomock.NewController(t)
	mockRepo := interfaces.NewMockProducts(ctrl)
	mockCache := interfaces.NewMockRedis(ctrl)
	productService := NewProductService(mockRepo, mockCache)

	return &ProductServiceTest{
		Controller: 	ctrl,
		MockRepo: 		mockRepo,
		MockCache: 		mockCache,
		ProductService: productService,
	}
}

// Helper function to create mock product data
func createMockProduct(id, sku, name, category string, price int) *models.Product {
	return &models.Product{
		ID: 			 id,
		Sku:       sku,
		Name:      name,
		Category:  category,
		Price:     price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Helper function to create expected response data
func createExpectedResponse(product *models.Product, finalPrice int, discount string) *models.ProductWithDiscountResponse {
	return &models.ProductWithDiscountResponse{
		ID: 			product.ID,
		Sku:      product.Sku,
		Name:     product.Name,
		Category: product.Category,
		Price: models.PriceWithDiscount{
			Original:           utils.ConvertPriceToDisplayFormat(product.Price),
			Final:              utils.ConvertPriceToDisplayFormat(finalPrice),
			DiscountPercentage: discount,
			Currency:           "EUR",
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}