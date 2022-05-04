package aggregate

// 商品聚合连接商品实体

import (
	"errors"
	"github.com/google/uuid"
	"github.com/wsqyouth/tech_note/ddd_go/entity"
)

var (
	// ErrMissingValues is returned when a product is created without a name or description
	ErrMissingValues = errors.New("missing values")
)

// Product 聚合商品实体、价格和数量
type Product struct {
	// item 是Porduct的根实体
	item     *entity.Item
	price    float64
	quantity int
}

// NewProduct 创建商品的工厂
func NewProduct(name, desc string, price float64) (Product, error) {
	// validate not empty
	if name == "" || desc == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: desc,
		},
		price:    price,
		quantity: 0,
	}, nil
}

/* 围绕customer product的贫血模型---start*/
func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}
func (p Product) GetPrice() float64 {
	return p.price
}

/* 围绕customer product的贫血模型---end*/
