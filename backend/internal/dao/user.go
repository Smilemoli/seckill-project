// file: backend/internal/dao/user.go
package dao

import (
	"errors"
	"seckill-project/backend/internal/model"

	"gorm.io/gorm"
)

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (bool, error) {
	var user model.User
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // 用户不存在，是正常情况
		}
		return false, err // 查询出错
	}
	return true, nil // 用户已存在
}

// CreateUser 在数据库中创建一个新用户
func CreateUser(user *model.User) error {
	return DB.Create(user).Error
}

// GetUserByUsername 根œ据用户名查询用户
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在，返回 nil, nil
		}
		return nil, err // 查询出错
	}
	return &user, nil
}
