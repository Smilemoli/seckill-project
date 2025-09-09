// file: backend/internal/model/activity.go
package model

import "time"

// SeckillActivity corresponds to the `seckill_activities` table
type SeckillActivity struct {
	BaseModel
	ProductID    uint      `gorm:"not null" json:"product_id"`
	SeckillPrice float64   `gorm:"type:decimal(10,2);not null" json:"seckill_price"`
	Stock        int       `gorm:"not null" json:"stock"`
	StartTime    time.Time `gorm:"not null" json:"start_time"`
	EndTime      time.Time `gorm:"not null" json:"end_time"`
}
