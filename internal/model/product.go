package model

import "time"

// Model utama di dalam DB
type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Name      string    `json:"name" example:"Kopi Kapal Api"`
	Price     float64   `json:"price" example:"17000"`
	Stock     int       `json:"stock" example:"13"`
	CreatedAt time.Time `json:"created_at" example:"0001-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"0001-01-01T00:00:00Z"`
}

// DTO yang digunakan untuk Request
type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required,min=2,max=100" example:"Kopi Kapal Api"`
	Price float64 `json:"price" binding:"required,gt=0" example:"1500.00"`
	Stock int     `json:"stock" binding:"required,gte=0" example:"15"`
}

// DTO khusus untuk Request Update
type UpdateProductRequest struct {
	Name  string  `json:"name" example:"Mechanical Keyboard Pro"`
	Price float64 `json:"price" example:"99.99"`
	Stock int     `json:"stock" example:"20"`
}

// Response Sukses
type ProductResponse struct {
	Message string  `json:"message" example:"success"`
	Data    Product `json:"data"`
}

// Response Error
type ErrorResponse struct {
	Error string `json:"error" example:"Produk tidak ditemukan!"`
}
