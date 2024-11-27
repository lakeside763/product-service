package interfaces

import "github.com/lakeside763/product-service/internal/core/models"


type Products interface {
	GetProducts(category string, priceLessThan int, cursorId string, pageSize int) ([]*models.Product, string, error)
	GetMaxDiscount(category string, sku string) (float64, error)
	CreateProduct(p models.CreateProductInput) (*models.Product, error)
	CheckProductExistsByName(name string) error
}