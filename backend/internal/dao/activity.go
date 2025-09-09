package dao

import (
	"seckill-project/backend/internal/model"
	"time"

	"gorm.io/gorm"
)

// CreateActivity 创建秒杀活动
func CreateActivity(activity *model.SeckillActivity) error {
	return DB.Create(activity).Error
}

// GetActivities 获取秒杀活动列表
func GetActivities() ([]model.SeckillActivity, error) {
	var activities []model.SeckillActivity

	err := DB.Where("end_time > ?", time.Now()).Find(&activities).Error
	return activities, err
}

// GetActivityByID 根据 ID 获取秒杀活动信息
func GetActivityByID(id uint) (*model.SeckillActivity, error) {
	var activity model.SeckillActivity
	err := DB.First(&activity, id).Error
	return &activity, err
}

// DecreaseActivityStock 扣减活动库存 (在事务中使用)
func DecreaseActivityStock(tx *gorm.DB, id uint, num int) error {
	// 使用 gorm 的表达式来执行 UPDATE ... SET stock = stock - num
	// 并增加条件 stock >= num 防止库存变为负数
	result := tx.Model(&model.SeckillActivity{}).Where("id = ? AND stock >= ?", id, num).Update("stock", gorm.Expr("stock - ?", num))
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
