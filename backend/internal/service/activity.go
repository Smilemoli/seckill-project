package service

import (
	"fmt"
	"seckill-project/backend/internal/dao"
	"seckill-project/backend/internal/model"
)

// CreateActivity 创建秒杀活动
func CreateActivity(activity *model.SeckillActivity) error {
	if err := dao.CreateActivity(activity); err != nil {
		return err
	}
	// 然后将库存预热到 Redis
	// key 的格式：seckill:activity:{id}:stock
	key := fmt.Sprintf("seckill:activity:%d:stock", activity.ID)
	return dao.RDB.Set(dao.Ctx, key, activity.Stock, 0).Err()
}

// GetActivities 获取秒杀活动列表
func GetActivities() ([]model.SeckillActivity, error) {
	return dao.GetActivities()
}
