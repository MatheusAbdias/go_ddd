package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/MatheusAbdias/go-ddd/aggregate"
	"github.com/MatheusAbdias/go-ddd/domain/customer"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}

	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrCustomerNotFound)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}
