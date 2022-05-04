package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/wsqyouth/tech_note/ddd_go/aggregate"
)

func Test_Tavern(t *testing.T) {
	// Create OrderService
	type testCase struct {
		name        string
		isUseMysql  bool
		expectedErr error
	}
	testCases := []testCase{
		// {
		// 	name:        "OrderService-memory",
		// 	expectedErr: nil,
		// },
		{
			name:        "OrderService-mysql",
			isUseMysql:  true,
			expectedErr: nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			products := init_products(t)
			var os *OrderService
			var err error
			if tt.isUseMysql {
				os, err = NewOrderService(
					WithMysqlCustomerRepository("coopers:2019Youth@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"),
					WithMemoryProductRepository(products),
				)
				if err != nil {
					t.Error(err)
				}
			} else {
				os, err = NewOrderService(
					WithMemoryCustomerRepository(),
					WithMemoryProductRepository(products),
				)
				if err != nil {
					t.Error(err)
				}
			}

			tavern, err := NewTavern(WithOrderService(os))
			if err != nil {
				t.Error(err)
			}

			cust, err := aggregate.NewCustomer("Coopers")
			if err != nil {
				t.Error(err)
			}

			err = os.customers.Add(cust)
			if err != nil {
				t.Error(err)
			}
			var productIDs []uuid.UUID
			for _, product := range products {
				productIDs = append(productIDs, product.GetID())
			}
			// Execute Order
			err = tavern.Order(cust.GetID(), productIDs)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
