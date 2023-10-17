package product

import (
	"errors"

	"github.com/MatheusAbdias/go-ddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound       = errors.New("no such product")
	ErrProductAllReadyExists = errors.New("there is already such an product")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetById(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
