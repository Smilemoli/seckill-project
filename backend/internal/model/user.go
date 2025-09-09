// file: backend/internal/model/user.go
package model

// User 对应于数据库中的 `users` 表
type User struct {
	BaseModel
	Username     string `gorm:"type:varchar(50);unique;not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
}
