package service

import (
	"log"

	"github.com/google/uuid"
)

type TarvernConfiguration func(os *Tavern) error

type Tavern struct {
	OrderService   *OrderService
	BillingService interface{}
}

func NewTaver(cfgs ...TarvernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithOrderSerivice(os *OrderService) TarvernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	log.Printf("\nBill the customer %v, order price:%v\n", customer, price)
	return nil
}
