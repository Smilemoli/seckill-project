// file: backend/internal/model/base.go
package model

import "time"

// BaseModel 定义了所有模型共有的字段
type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
