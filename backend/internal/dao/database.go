package dao

import (
	"seckill-project/backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() (err error) {
	dsn := config.Conf.MySQL.DSN

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	// 检查数据库连接
	sqlDB, err := DB.DB()
	if err != nil {
		return
	}

	return sqlDB.Ping()
}
