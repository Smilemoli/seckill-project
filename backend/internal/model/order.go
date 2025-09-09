// file: backend/internal/model/order.go
package model

// Order 对应于 `orders` 表
type Order struct {
	BaseModel
	UserID     uint    `gorm:"not null"`
	ActivityID uint    `gorm:"not null"`
	ProductID  uint    `gorm:"not null"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null"`
	Status     int8    `gorm:"type:tinyint;not null;default:1"` // 1-处理中, 2-成功, 3-失败
}
