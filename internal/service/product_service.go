package service

import (
	"go_tutor/internal/model"
	"go_tutor/internal/repository"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func (s *productService) CreateProduct(p *model.Product) error {
	return s.repo.Create(p)
}

func (s *productService) GetAllProducts() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetProductByID(id uint) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) UpdateProduct(p *model.Product) error {
	return s.repo.Update(p)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}