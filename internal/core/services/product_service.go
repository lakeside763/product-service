package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/lakeside763/product-service/internal/core/models"
	"github.com/lakeside763/product-service/internal/ports/interfaces"
	"github.com/lakeside763/product-service/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type ProductService struct {
	repo  interfaces.Products
	cache interfaces.Redis
}

func NewProductService(repo interfaces.Products, cache interfaces.Redis) *ProductService {
	return &ProductService{
		repo:  repo,
		cache: cache,
	}
}

func (s *ProductService) GetProductsWithDiscount(
	category string, priceLessThan int, cursorId string, pageSize int,
) ([]*models.ProductWithDiscountResponse, string, error) {
	// Get product passed on priceLessThan, lastProductId and pageSize
	products, nextCursorId, err := s.repo.GetProducts(category, priceLessThan, cursorId, pageSize)
	if err != nil {
		return nil, "", err
	}

	// Map products to ProductWithDiscountResponse
	var data []*models.ProductWithDiscountResponse
	for _, product := range products {
		// Apply discount to the product
		discount, err := s.getDiscountForProduct(product)
		if err != nil {
			return nil, "", err
		}
		data = append(data, s.mapToProductWithDiscountResponse(product, discount))
	}

	return data, nextCursorId, nil
}

func (s *ProductService) getDiscountForProduct(product *models.Product) (float64, error) {
	var discount float64
	cacheKey := fmt.Sprintf("discount-%s-%s", product.Category, product.Sku)
	cacheExpiration := 3 * 24 * time.Hour // Discount expires in 3 days

	// Check cache for an existing discount
	cachedDiscount, err := s.cache.Get(cacheKey)
	if err == nil && cachedDiscount != "" {
		if discount, err = strconv.ParseFloat(cachedDiscount, 64); err == nil {
			return discount, nil
		}
	}

	discount, err = s.repo.GetMaxDiscount(product.Category, product.Sku)
	if err != nil {
		log.Errorf("Failed to get discount from database: %v", err)
	}

	if err := s.cache.Set(cacheKey, fmt.Sprintf("%f", discount), cacheExpiration); err != nil {
		log.Errorf("Failed to save discount to cache")
	}

	return discount, nil
}

func (s *ProductService) mapToProductWithDiscountResponse(
	product *models.Product, discount float64,
) *models.ProductWithDiscountResponse {
	finalPrice := product.Price
	discountPercentage := ""
	if discount > 0 {
		discountPercentage = fmt.Sprintf("%.0f%%", discount)
		finalPrice = int(float64(product.Price) * (1 - discount/100))
	}

	return &models.ProductWithDiscountResponse{
		ID:       product.ID,
		Sku:      product.Sku,
		Name:     product.Name,
		Category: product.Category,
		Price: models.PriceWithDiscount{
			Original:           utils.ConvertPriceToDisplayFormat(product.Price),
			Final:              utils.ConvertPriceToDisplayFormat(finalPrice),
			DiscountPercentage: discountPercentage,
			Currency:           "EUR",
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
