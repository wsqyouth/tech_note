package valueobject

import (
	"github.com/google/uuid"
	"time"
)

// Transaction 交易订单, 这里模拟为值对象
type Transaction struct {
	// 值对象是immutable的，因此设置为lowercase
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
