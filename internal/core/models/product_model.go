package models

import "time"

type Product struct {
	ID				string			`json:"id" gorm:"type:int; column:id"`
	Sku       string			`json:"sku" gorm:"size:255; column:sku"`
	Name      string			`json:"name" gorm:"size:255; column:name"`
	Category  string			`json:"category" gorm:"size:255; column:category"`
	Price     int					`json:"price" gorm:"type:int; column:price"`
	CreatedAt time.Time		`json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time		`json:"updated_at" gorm:"column:updated_at"`
}

type PriceWithDiscount struct {
	Original           float64
	Final              float64
	DiscountPercentage string
	Currency           string
}

type ProductWithDiscountResponse struct {
	Sku string
	Name string
	Category string
	Price PriceWithDiscount
	CreatedAt time.Time
	UpdatedAt time.Time
}
