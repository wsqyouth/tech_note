package customer

// 仓储模式-Customer 定义repository所需的函数: Get,Add,Update
/*
the repository only handles the Customer aggregate and it should only do so.
Never make the repository coupled to any other aggregate, we want louse coupling.
*/

import (
	"errors"
	"github.com/google/uuid"
	"github.com/wsqyouth/tech_note/ddd_go/aggregate"
)

var (
	// ErrCustomerNotFound is returned when a customer is not found.
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")
	// ErrFailedToAddCustomer is returned when the customer could not be added to the repository.
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	// ErrUpdateCustomer is returned when the customer could not be updated in the repository.
	ErrUpdateCustomer = errors.New("failed to update the customer in the repository")
)

// CustomerRepository接口定义行为规则
type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
