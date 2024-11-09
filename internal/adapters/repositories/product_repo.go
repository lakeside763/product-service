package repositories

import (
	"github.com/lakeside763/product-service/internal/core/models"
	"github.com/lakeside763/product-service/internal/ports/interfaces"
	"github.com/lakeside763/product-service/pkg/utils"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) interfaces.Products {
	return &ProductRepo{DB: db}
}

// GetProducts implements interfaces.Products.
func (repo *ProductRepo) GetProducts(priceLessThan int, lastProductId string, pageSize int) ([]*models.Product, error) {
	var products []*models.Product

	pageSize = max(pageSize, 10)
	priceLessThan = utils.ConvertPriceToStoredFormat(priceLessThan)

	// Start the query
	query := repo.DB.Limit(pageSize).Order("id ASC")

	// Apply the price filter conditionally
	if priceLessThan > 0 {
		query = query.Where("price < ?", priceLessThan)
	}

	// Apply the ID filter only if lastProductId is not empty
	if lastProductId != "" {
		query = query.Where("id > ?", lastProductId)
	}

	// Execute the query
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepo) GetMaxDiscount(category string, sku string) (float64, error) {
	var maxDiscount float64
	err := repo.DB.Table("discounts").
				Select("MAX(discount_percentage)").
				Where("category = ? OR sku = ?", category, sku).
				Scan(&maxDiscount).Error

	if err != nil {
		return 0, err
	}

	return maxDiscount, nil
}



