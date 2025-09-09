// file: backend/internal/model/product.go
package model

// Product 对应于数据库中的 `products` 表
type Product struct {
	BaseModel
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Stock       int     `gorm:"not null;default:0"`
}
