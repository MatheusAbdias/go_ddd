package aggregate

import (
	"errors"

	"github.com/google/uuid"

	"github.com/MatheusAbdias/go-ddd/entity"
	"github.com/MatheusAbdias/go-ddd/valueobject"
)

var (
	ErrInvalidPerson = errors.New("a customer must be has valid name")
)

type Customer struct {
	person      *entity.Person
	products    []*entity.Item
	transaction []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
			person:      person,
			products:    make([]*entity.Item, 0),
			transaction: make([]valueobject.Transaction, 0),
		},
		nil

}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
