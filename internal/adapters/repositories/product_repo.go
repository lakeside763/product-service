package repositories

import (
	"fmt"

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
func (repo *ProductRepo) GetProducts(category string, priceLessThan int, cursorId string, pageSize int) ([]*models.Product, string, error) {
	var products []*models.Product

	pageSize = utils.PageSize(pageSize)
	priceLessThan = utils.ConvertPriceToStoredFormat(priceLessThan)

	// Start the query
	query := repo.DB.Where("category = ?", category).Limit(pageSize).Order("serial_id ASC")

	// Apply the price filter conditionally
	if priceLessThan > 0 {
		query = query.Where("price < ?", priceLessThan)
	}

	// Apply the ID filter only if nextCursorId is not empty
	if cursorId != "" {
		serialId, err := utils.DecodeCursorId(cursorId)
		if err != nil {
			return nil, "", fmt.Errorf("invalid cursor: %w", err)
		}
		query = query.Where("serial_id > ?", serialId)
	}

	// Execute the query
	if err := query.Find(&products).Error; err != nil {
		return nil, "", err
	}

	var nextCursorId string
	if len(products) > 0 {
		lastProduct := products[len(products)-1]
		nextCursorId = utils.EncodeCursorId(lastProduct.SerialId)
	}

	return products, nextCursorId, nil
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



