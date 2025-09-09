// file: backend/internal/dao/order.go
package dao

import (
	"seckill-project/backend/internal/model"

	"gorm.io/gorm"
)

// CreateOrder 创建订单 (在事务中使用)
func CreateOrder(tx *gorm.DB, order *model.Order) error {
	return tx.Create(order).Error
}
