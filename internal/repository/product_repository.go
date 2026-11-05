package repository

import (
	"go_tutor/internal/model"
	"sync"
)

type ProductRepository interface {
	Create(product *model.Product) error
	FindAll() ([]model.Product, error)
	FindByID(id uint) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uint) error
}

type memoryRepository struct {
	mu       sync.RWMutex
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
	r.mu.Lock()
	defer r.mu.Unlock()

	p.ID = r.nextID
	r.products[p.ID] = *p
	r.nextID++
	return nil
}

func (r *memoryRepository) FindAll() ([]model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]model.Product, 0, len(r.products))
	for _, p := range r.products {
		list = append(list, p)
	}
	return list, nil
}

func (r *memoryRepository) FindByID(id uint) (*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, exists := r.products[id]
	if !exists {
		return nil, model.ErrProductNotFound
	}
	return &p, nil
}

func (r *memoryRepository) Update(p *model.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[p.ID]; !exists {
		return model.ErrProductNotFound
	}
	r.products[p.ID] = *p
	return nil
}

func (r *memoryRepository) Delete(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[id]; !exists {
		return model.ErrProductNotFound
	}
	delete(r.products, id)
	return nil
}
