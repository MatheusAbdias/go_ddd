package memory

import (
	"sync"

	"github.com/MatheusAbdias/go-ddd/aggregate"
	"github.com/MatheusAbdias/go-ddd/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, produt := range mpr.products {
		products = append(products, produt)
	}
	return products, nil

}

func (mpr *MemoryProductRepository) GetById(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}
func (mpr *MemoryProductRepository) Add(newProd aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newProd.GetID()]; ok {
		return product.ErrProductAllReadyExists
	}
	mpr.products[newProd.GetID()] = newProd
	return nil
}

func (mpr *MemoryProductRepository) Update(update aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[update.GetID()]; !ok {
		return product.ErrProductNotFound
	}
	mpr.products[update.GetID()] = update
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(mpr.products, id)

	return nil
}
