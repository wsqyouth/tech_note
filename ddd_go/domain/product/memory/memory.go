package memory

// product-memory是product repository的内存实现

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/wsqyouth/tech_note/ddd_go/aggregate"
	"github.com/wsqyouth/tech_note/ddd_go/domain/product"
	"sync"
)

// MemoryProductRepository 实现接口
type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

// New 生成新的 product repositroy
func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	//map -> slice
	var products []aggregate.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if _, ok := mpr.products[id]; ok {
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

// Add 新增
func (mpr *MemoryProductRepository) Add(product aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[product.GetID()]; ok {
		// return fmt.Errorf("product already exist: %w", product.ErrProductAlreadyExist)
		return product.ErrProductAlreadyExist
	}
	mpr.products[product.GetID()] = product

	return nil
}

// Update 更新
func (mpr *MemoryProductRepository) Update(product aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[product.GetID()]; !ok {
		// return fmt.Errorf("product does not exist: %w", product.ErrProductNotFound)
		return product.ErrProductNotFound
	}

	mpr.products[product.GetID()] = product
	return nil
}

// Delete 删除
func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[id]; !ok {
		return fmt.Errorf("product does not exist: %w", product.ErrProductNotFound)
	}
	delete(mpr.products, id)
	return nil
}
