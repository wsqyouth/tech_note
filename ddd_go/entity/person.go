package entity

import (
	"github.com/google/uuid"
)

// Person 领域内实体
type Person struct {
	// id是实体的标识，被所有子域访问
	ID   uuid.UUID
	Name string
	Age  int
}
