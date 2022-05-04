package aggregate

// 聚合连接实体和值对象

import (
	"errors"
	"github.com/google/uuid"
	"github.com/wsqyouth/tech_note/ddd_go/entity"
	"github.com/wsqyouth/tech_note/ddd_go/valueobject"
)

// Customer is a aggregate that combines all entities needed
type Customer struct {
	// person是custome的根实体,即意味着persion.ID是这个聚合的main identifier
	person *entity.Person
	// 一个customer可以有很多products
	products []*entity.Item
	// 一个custome可以有很多transactions
	transactions []valueobject.Transaction
}

var (
	// ErrInvalidPerson is returned when the person is not valid in the NewCustome factory
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

// NewCustomer 创造new custome的工厂
func NewCustomer(name string) (Customer, error) {
	// validate name not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	// 创建customer对象并初始化成员不为nil pointer
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

/* 围绕customer persion的贫血模型---start*/
// GetID 获取customer的根实体id
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID 设置根实体id
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

// GetName 获取customer的name
func (c *Customer) GetName() string {
	return c.person.Name
}

// SetName 设置根实体的name
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

/* 围绕customer persion的贫血模型---end*/
