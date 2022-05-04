package services

// loosely coupled repositories into a business flow以实现特定领域需要
import (
	"context"
	"github.com/google/uuid"
	"github.com/wsqyouth/tech_note/ddd_go/aggregate"
	"github.com/wsqyouth/tech_note/ddd_go/domain/customer"
	"github.com/wsqyouth/tech_note/ddd_go/domain/customer/memory"
	"github.com/wsqyouth/tech_note/ddd_go/domain/customer/mysql_custom"
	"github.com/wsqyouth/tech_note/ddd_go/domain/product"
	prodmemory "github.com/wsqyouth/tech_note/ddd_go/domain/product/memory"
	"log"
)

// OrderService implementation
type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

// neat trick
type OrderConfiguration func(os *OrderService) error

// NewOrderService 工厂模式: 输入多个OrderConfiguration,处理后生成new OrderService
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// 新orderservice
	os := &OrderService{}
	// Apply多个OrderConfiguration
	for _, cfg := range cfgs {
		if err := cfg(os); err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository [接口]将customer repository设置为OrderConfiguration
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository [实现]将特定customer memory repository设置为OrderConfiguration
func WithMemoryCustomerRepository() OrderConfiguration {
	//创建customer memory repo,如果需要添加参数在这里设置
	cr := memory.New()
	return WithCustomerRepository(cr)
}

// WithMemoryProductRepository [实现]将特定product memory repository设置为OrderConfiguration
func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// 创建product memory repo,如果需要添加参数在这里设置
		pr := prodmemory.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}

}

// CreateOrder 将多个repositories组合起来实现订单功能,返回商品总价钱
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// 1. 获取customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	// 2. 获取商品
	var products []aggregate.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}
	log.Printf("Customer: %s has order %d products. price:%v", c.GetID(), len(products), price)
	return price, nil
}

// WithMysqlCustomerRepository mysql存储customer
func WithMysqlCustomerRepository(connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		//创建customer mysql repo,如果需要添加参数在这里设置
		cr, err := mysql_custom.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}
