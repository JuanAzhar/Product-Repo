package entity

import (
	"time"
)

type ProductsCore struct {
	ID            string    `json:"id"`
	Product_image string    `json:"product_image"`
	Product_name  string    `json:"product_name"`
	Price         string    `json:"price"`
	Quantity      int       `json:"quantity"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"update_at"`
}
