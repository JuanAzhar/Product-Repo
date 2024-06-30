package model

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID         uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	Username   string    `gorm:"varchar(50);not null" json:"username"`
	Email      string    `gorm:"varchar(50);not null" json:"email"`
	Password   string    `gorm:"varchar(50);not null" json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"update_at"`
}
