package aggregate

import (
	"errors"

	"github.com/google/uuid"

	"github.com/MatheusAbdias/go-ddd/entity"
)

var (
	ErrMissingValues = errors.New("missing importante values")
)

type Product struct {
	item     *entity.Item
	price    int
	quantity int
}

func NewProduct(name, description string, price int) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() int {
	return p.price
}
