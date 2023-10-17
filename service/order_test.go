package service

import (
	"testing"

	"github.com/google/uuid"

	"github.com/MatheusAbdias/go-ddd/aggregate"
)

func init_produts(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy", 199)
	if err != nil {
		t.Fatal(err)
	}

	peenuts, err := aggregate.NewProduct("Peenuts", "Snacks", 99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "nasty drink", 999)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		beer, peenuts, wine,
	}

}

func TestOrderService(t *testing.T) {
	products := init_produts(t)

	os, err := NewOrderService(
		WithMemoryCustomerRespository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	customer, err := aggregate.NewCustomer("Mat")
	if err != nil {
		t.Fatal(err)
	}

	err = os.customers.Add(customer)
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}
	_, err = os.CreateOrder(customer.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}
