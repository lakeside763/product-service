package interfaces

import "github.com/lakeside763/product-service/internal/core/models"


type Products interface {
	GetProducts() ([]*models.Product, error)
}