package customer

import (
	"errors"

	"github.com/google/uuid"

	"github.com/MatheusAbdias/go-ddd/aggregate"
)

var (
	ErrCustomerNotFound     = errors.New("customer not found in repository")
	ErrFailedToAddCustomer  = errors.New("failed to add  the customer")
	ErrFailToUpdateCustomer = errors.New("afiled to update the customer")
)

type CustomerRespository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
