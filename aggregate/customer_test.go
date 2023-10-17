package aggregate_test

import (
	"errors"
	"testing"

	"github.com/MatheusAbdias/go-ddd/aggregate"
)

func TestCustomerNewCustomer(t *testing.T) {
	type TestCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []TestCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "Percy Bolmer",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
