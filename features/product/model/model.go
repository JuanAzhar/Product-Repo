package model

import (
	"time"

	"github.com/google/uuid"
)

type Products struct {
	ID            uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	Product_image string    `json:"product_image"`
	Product_name  string    `json:"Product_name"`
	Price         string    `json:"price"`
	Quantity      int       `json:"quantity"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"update_at"`
}
