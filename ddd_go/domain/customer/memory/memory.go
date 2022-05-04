package memory

// memeory 是customer repository的内存实现

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/wsqyouth/tech_note/ddd_go/aggregate"
	"github.com/wsqyouth/tech_note/ddd_go/domain/customer"
	"sync"
)

// MemoryRepository 实现接口
type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// New 生成新的 customer repositroy
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get 根据id查找
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound //妙,错误定义在接口处
}

// Add 新增
func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		// Saftey check
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exist: %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update 更新
func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
