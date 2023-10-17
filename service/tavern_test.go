package service

import (
	"testing"

	"github.com/google/uuid"

	"github.com/MatheusAbdias/go-ddd/aggregate"
)

func Test_Tavern(t *testing.T) {
	products := init_produts(t)

	os, err := NewOrderService(
		WithMemoryCustomerRespository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTaver(WithOrderSerivice(os))
	if err != nil {
		t.Fatal(err)
	}

	customer, err := aggregate.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}
	if err = os.customers.Add(customer); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(customer.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}

}
