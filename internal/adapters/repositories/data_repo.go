package repositories

import (
	"github.com/lakeside763/product-service/internal/adapters/database"
	"github.com/lakeside763/product-service/internal/ports/interfaces"
)

type DataRepo struct {
	Product interfaces.Products
}

func NewDataRepo(databaseUrl string) (*DataRepo, error) {
	db, err := database.PostgresDB(databaseUrl)
	if err != nil {
		return nil, err
	}
	
	productRepo := NewProductRepo(db)

	return &DataRepo{
		Product: productRepo,
	}, nil
}