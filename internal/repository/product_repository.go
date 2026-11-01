package repository

import (

	"errors"
	"go_tutor/internal/model"
)

type ProductRepository interface {
	Create(product *model.Product) error
	FindAll() ([]model.Product, error)
	FindByID(id uint) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uint) error
}

type memoryRepository struct {
	products map[uint]model.Product
	nextID   uint
}

func NewMemoryRepository() ProductRepository {
	return &memoryRepository{
		products: make(map[uint]model.Product),
		nextID:   1,
	}
}

func (r *memoryRepository) Create(p *model.Product) error {
	p.ID = r.nextID
	r.products[p.ID] = *p
	r.nextID++
	return nil
}

func (r *memoryRepository) FindAll() ([]model.Product, error) {
	list := make([]model.Product, 0, len(r.products))
	for _, p := range r.products {
		list = append(list, p)
	}
	return list, nil
}

func (r *memoryRepository) FindByID(id uint) (*model.Product, error) {
	p, exists := r.products[id]
	if !exists {
		return nil, errors.New("Product not found")
	}
	return &p, nil
}

func (r *memoryRepository) Update(p *model.Product) error {
	if _, exists := r.products[p.ID]; !exists {
		return errors.New("Product not found")
	}
	r.products[p.ID] = *p
	return nil
}

func (r *memoryRepository) Delete(id uint) error {
	if _, exists := r.products[id]; !exists {
		return errors.New("Product not found")
	}
	delete(r.products, id)
	return nil
}