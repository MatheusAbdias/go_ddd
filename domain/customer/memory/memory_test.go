package memory

import (
	"errors"
	"testing"

	"github.com/google/uuid"

	"github.com/MatheusAbdias/go-ddd/aggregate"
	"github.com/MatheusAbdias/go-ddd/domain/customer"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name     string
		id       uuid.UUID
		expected error
	}

	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:     "no customer by id",
			id:       uuid.New(),
			expected: customer.ErrCustomerNotFound,
		},
		{
			name:     "customer by id",
			id:       id,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expected) {
				t.Errorf("Expected error %v,got %v", tc.expected, err)
			}
		})
	}
}
