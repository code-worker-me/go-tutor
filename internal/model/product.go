package model

import "time"

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required,min=2,max=100"`
	Price     float64   `json:"price" binding:"required,gt=0"`
	Stock     int       `json:"stock" binding:"required,gte=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
