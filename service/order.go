package service

import (
	"log"

	"github.com/google/uuid"

	"github.com/MatheusAbdias/go-ddd/aggregate"
	"github.com/MatheusAbdias/go-ddd/domain/customer"
	"github.com/MatheusAbdias/go-ddd/domain/customer/memory"
	"github.com/MatheusAbdias/go-ddd/domain/product"
	prodmem "github.com/MatheusAbdias/go-ddd/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRespository
	prodcuts  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRespository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRespository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}

		os.prodcuts = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (int, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var total int

	for _, id := range productsIDs {
		p, err := o.prodcuts.GetById(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has order %d products", c.GetID(), len(products))

	return total, nil
}
