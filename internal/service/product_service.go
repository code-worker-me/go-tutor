package service

import (
	"errors"
	"go_tutor/internal/model"
	"go_tutor/internal/repository"
	"time"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	UpdateProduct(req *model.UpdateProductRequest, id uint) (*model.Product, error)
	DeleteProduct(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
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

func (s *productService) UpdateProduct(req *model.UpdateProductRequest, id uint) (*model.Product, error) {
	existingProduct, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("Product not found")
	}

	isChanged := false

	if req.Name != "" && req.Name != existingProduct.Name {
		existingProduct.Name = req.Name
		isChanged = true
	}
	if req.Price != 0 && req.Price != existingProduct.Price {
		existingProduct.Stock = req.Stock
		isChanged = true
	}
	if req.Stock != 0 && req.Stock != existingProduct.Stock {
		existingProduct.Stock = req.Stock
		isChanged = true
	}

	if isChanged {
		existingProduct.UpdatedAt = time.Now()
		if err := s.repo.Update(existingProduct); err != nil {
			return nil, err
		}
	}

	return existingProduct, nil
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}