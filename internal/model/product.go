package model

import "time"

// Model utama di dalam DB
type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DTO yang digunakan untuk Request
type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required,min=2,max=100" example:"Kopi Kapal Api"`
	Price float64 `json:"price" binding:"required,gt=0" example:"1500.00"`
	Stock int     `json:"stock" binding:"required,gte=0" example:"15"`
}
