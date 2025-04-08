package model

import (
	"time"
)

const (
	StatusInStock    = "In Stock"
	StatusOutOfStock = "Out of stock"
)

type Product struct {
	ID            int64     `gorm:"primaryKey" json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	StockQuantity int64     `json:"stock_quantity"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
