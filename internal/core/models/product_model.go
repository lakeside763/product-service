package models

import (
	"time"
)

type Product struct {
	ID        string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	SerialId  int64     `json:"serial_id" gorm:"type:int;column:serial_id;<-:false"`
	Sku       string    `json:"sku" gorm:"size:255; column:sku;<-:false"`
	Name      string    `json:"name" gorm:"size:255; column:name"`
	Category  string    `json:"category" gorm:"size:255; column:category"`
	Price     int       `json:"price" gorm:"type:int; column:price"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type PriceWithDiscount struct {
	Original           float64 `json:"original"`
	Final              float64 `json:"final"`
	DiscountPercentage string  `json:"discount_percentage"`
	Currency           string  `json:"currency"`
}

type ProductWithDiscountResponse struct {
	ID        string            `json:"id"`
	Sku       string            `json:"sku"`
	Name      string            `json:"name"`
	Category  string            `json:"category"`
	Price     PriceWithDiscount `json:"price"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type Discount struct {
	ID                 int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Sku                string    `json:"sku" gorm:"size:255"`
	Category           string    `json:"category" gorm:"size:255"`
	DiscountPercentage float64   `json:"discount_percentage" gorm:"type:decimal(5,2)"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type ProductDataResponse struct {
	Data     []*ProductWithDiscountResponse `json:"data"`
	CursorId string                         `json:"cursorId"`
}

type CreateProductInput struct {
	Name string				`json:"name"`
	Category string		`json:"category"`
	Price float64			`json:"price"`
}
