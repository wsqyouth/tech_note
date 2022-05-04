package services

import (
	"log"

	"github.com/google/uuid"
)

// Tavern 是整个服务入口,将OrderService作为sub-service
type Tavern struct {
	// orderService 处理订单服务
	OrderService *OrderService
	// BillingService 处理余额服务,取决于你的实现
	BillingService interface{}
}

// neat trick
type TavernConfiguration func(t *Tavern) error

// NewTavern
func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	//新Tavern
	tavern := &Tavern{}
	// Apply多个TavernConfiguration
	for _, cfg := range cfgs {
		if err := cfg(tavern); err != nil {
			return nil, err
		}
	}
	return tavern, nil
}

// WithOrderService 将特定orderService设置为TavernConfiguration
func WithOrderService(os *OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

// Order 创建订单服务
func (t *Tavern) Order(customerID uuid.UUID, productIDs []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customerID, productIDs)
	if err != nil {
		return err
	}
	log.Printf("Order price :%v", price)

	// Bill the customer
	// err = t.BillingService.Bill(customer,err)
	return nil
}
