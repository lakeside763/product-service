package repositories

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/lakeside763/product-service/internal/core/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProductRepo_GetProducts(t *testing.T) {
	db, repo := setupTestDB(t)

	// Insert sample products
	products := []models.Product{
		newProduct("000001", "BV Lean leather ankle boots", "boots", 10000),
		newProduct("000002", "BV Lean leather ankle boots", "boots", 20000),
		newProduct("000003", "Naima embellished suede sandals", "sandals", 30000),
	}
	db.Create(&products)
	

	// Test case 1: Retrieve products with no price filter
	result, err := repo.GetProducts("boots", 0, "", 10)
	assert.NoError(t, err)
	assert.Len(t, result, 2) 
	assert.Equal(t, "BV Lean leather ankle boots", result[0].Name)

	// Test case 2: Retrieve products with priceLessThan filter
	result, err = repo.GetProducts("boots", 200, "", 10)
	assert.NoError(t, err)
	assert.Len(t, result, 1) // Expecting two products priced below 2500
}

func TestProductRepo_GetMaxDiscount(t *testing.T) {
	db, repo := setupTestDB(t)

	// Insert sample discounts
	discounts := []models.Discount{
		{Category: "boots", Sku: "000001", DiscountPercentage: 15.0},
		{Category: "boots", Sku: "000002", DiscountPercentage: 20.0},
		{Category: "sandals", Sku: "000003", DiscountPercentage: 10.0},
	}
	db.Create(&discounts)

	// Test case 1: Retrieve max discount for category "Electronics"
	maxDiscount, err := repo.GetMaxDiscount("boots", "")
	assert.NoError(t, err)
	assert.Equal(t, 20.0, maxDiscount)

	// Test case 2: Retrieve max discount for SKU "003"
	maxDiscount, err = repo.GetMaxDiscount("", "000002")
	assert.NoError(t, err)
	assert.Equal(t, 20.0, maxDiscount)

	// Test case 3: Retrieve max discount for both category "Apparel" and SKU "001"
	maxDiscount, err = repo.GetMaxDiscount("boots", "000001")
	assert.NoError(t, err)
	assert.Equal(t, 20.0, maxDiscount) // Assuming highest between category and SKU is taken
}

func setupTestDB(t *testing.T) (*gorm.DB, *ProductRepo) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	// Migrate models to create the tables
	if err := db.AutoMigrate(&models.Product{}, &models.Discount{}); err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}

	return db, NewProductRepo(db).(*ProductRepo)
}

func newProduct(sku, name, category string, price int) models.Product {
	return models.Product{
		ID:        uuid.New().String(),
		Sku:       sku,
		Name:      name,
		Category:  category,
		Price:     price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
