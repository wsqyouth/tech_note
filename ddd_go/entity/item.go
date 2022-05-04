package entity

import (
	"github.com/google/uuid"
)

// Item 领域内实体
type Item struct {
	// id是实体的标识，被所有子域访问
	ID          uuid.UUID
	Name        string
	Description string
}
