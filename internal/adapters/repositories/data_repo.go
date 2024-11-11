package repositories

import (
	"github.com/lakeside763/product-service/internal/adapters/database"
	"github.com/lakeside763/product-service/internal/ports/interfaces"
	"gorm.io/gorm"
)

type DataRepo struct {
	Product interfaces.Products
	db      *gorm.DB
}

func NewDataRepo(databaseUrl string) (*DataRepo, error) {
	db, err := database.PostgresDB(databaseUrl)
	if err != nil {
		return nil, err
	}
	
	productRepo := NewProductRepo(db)

	return &DataRepo{
		Product: productRepo,
		db:      db,
	}, nil
}

// Close method to close the database connection
func (repo *DataRepo) Close() error {
	sqlDB, err := repo.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}