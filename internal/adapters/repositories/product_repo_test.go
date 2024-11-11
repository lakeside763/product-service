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
		newProduct("001", "Product1", "Electronics", 10000),
		newProduct("002", "Product2", "Electronics", 20000),
		newProduct("003", "Product3", "Apparel", 30000),
	}
	db.Create(&products)
	

	// Test case 1: Retrieve products with no price filter
	result, err := repo.GetProducts(0, "", 10)
	assert.NoError(t, err)
	assert.Len(t, result, 3) 
	assert.Equal(t, "Product1", result[0].Name)

	// Test case 2: Retrieve products with priceLessThan filter
	result, err = repo.GetProducts(250, "", 10)
	assert.NoError(t, err)
	assert.Len(t, result, 2) // Expecting two products priced below 2500
}

func TestProductRepo_GetMaxDiscount(t *testing.T) {
	db, repo := setupTestDB(t)

	// Insert sample discounts
	discounts := []models.Discount{
		{Category: "Electronics", Sku: "001", DiscountPercentage: 15.0},
		{Category: "Electronics", Sku: "002", DiscountPercentage: 20.0},
		{Category: "Apparel", Sku: "003", DiscountPercentage: 10.0},
	}
	db.Create(&discounts)

	// Test case 1: Retrieve max discount for category "Electronics"
	maxDiscount, err := repo.GetMaxDiscount("Electronics", "")
	assert.NoError(t, err)
	assert.Equal(t, 20.0, maxDiscount)

	// Test case 2: Retrieve max discount for SKU "003"
	maxDiscount, err = repo.GetMaxDiscount("", "003")
	assert.NoError(t, err)
	assert.Equal(t, 10.0, maxDiscount)

	// Test case 3: Retrieve max discount for both category "Apparel" and SKU "001"
	maxDiscount, err = repo.GetMaxDiscount("Apparel", "001")
	assert.NoError(t, err)
	assert.Equal(t, 15.0, maxDiscount) // Assuming highest between category and SKU is taken
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
