package repositories

import (
	"github.com/lakeside763/product-service/internal/core/models"
	"github.com/lakeside763/product-service/internal/ports/interfaces"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) interfaces.Products {
	return &ProductRepo{db: db}
}

// GetProducts implements interfaces.Products.
func (repo *ProductRepo) GetProducts() ([]*models.Product, error) {
	var products []*models.Product
	if err := repo.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}


