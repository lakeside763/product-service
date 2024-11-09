package services

import (
	"github.com/lakeside763/product-service/internal/core/models"
	"github.com/lakeside763/product-service/internal/ports/interfaces"
)

type ProductService struct {
	repo interfaces.Products
}

func NewProductService(repo interfaces.Products) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]*models.Product, error) {
	// data := []models.Product{
	// 	{Name: "Jacket", Price: 100},
	// 	{Name: "Boot", Price: 200},
	// }

	return s.repo.GetProducts()
	
}