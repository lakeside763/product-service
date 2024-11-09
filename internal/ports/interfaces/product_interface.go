package interfaces

import "github.com/lakeside763/product-service/internal/core/models"


type Products interface {
	GetProducts(priceLessThan int, lastProductId string, pageSize int) ([]*models.Product, error)
	GetMaxDiscount(category string, sku string) (float64, error)
}