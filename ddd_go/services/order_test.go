package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/wsqyouth/tech_note/ddd_go/aggregate"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("啤酒", "青岛", 100)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := aggregate.NewProduct("红酒", "法国", 200)
	if err != nil {
		t.Error(err)
	}
	wine, err := aggregate.NewProduct("白酒", "中国", 300)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{
		beer, peenuts, wine,
	}
	return products
}
func TestOrder_NewOrderService(t *testing.T) {
	// Create a few products to insert into in memory repo
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}

	// Add Customer
	cust, err := aggregate.NewCustomer("coopers")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	// Perform Order for all beer
	var productIDs []uuid.UUID
	for _, product := range products {
		productIDs = append(productIDs, product.GetID())
	}

	_, err = os.CreateOrder(cust.GetID(), productIDs)

	if err != nil {
		t.Error(err)
	}

}
