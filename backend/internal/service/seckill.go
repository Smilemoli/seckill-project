// file: backend/internal/service/seckill.go
package service

import (
	"errors"
	"fmt"
	"seckill-project/backend/internal/dao"
	"time"

	"gorm.io/gorm"
)

// ExecuteSeckillHighConcurrency 是高并发改造后的秒杀核心逻辑
// 它不再返回*model.Order，因为订单是异步创建的
func ExecuteSeckillHighConcurrency(userID, activityID uint) error {
	// 1. 数据校验：活动是否存在、是否在时间范围内
	activity, err := dao.GetActivityByID(activityID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("秒杀活动不存在")
		}
		return err // 其他数据库错误
	}

	now := time.Now()
	if now.Before(activity.StartTime) {
		return errors.New("秒杀尚未开始")
	}
	if now.After(activity.EndTime) {
		return errors.New("秒杀已经结束")
	}

	// 2. Redis 预减库存
	// 构造 Redis 键
	key := fmt.Sprintf("seckill:activity:%d:stock", activityID)

	// 执行 DECR 原子减一操作
	stock, err := dao.RDB.Decr(dao.Ctx, key).Result()
	if err != nil {
		// 如果 Redis 出错，应该记录日志并快速失败
		return errors.New("系统繁忙，请稍后再试 (Redis error)")
	}

	// 3. 检查库存结果
	if stock < 0 {
		// 库存不足，说明在 DECR 之前库存已经是 0 了
		// 为了防止负数库存持续减少，需要把减掉的库存加回来
		dao.RDB.Incr(dao.Ctx, key)
		return errors.New("商品已售罄")
	}

	// 4. 异步下单：发送消息到 RabbitMQ
	// 库存预减成功，说明抢到了资格
	msg := OrderMessage{
		UserID:     userID,
		ActivityID: activityID,
	}

	// 使用 goroutine 异步发送消息，这样不会阻塞当前秒杀请求的响应
	go PublishOrderMessage(msg)

	// 5. 直接返回成功
	// 告诉用户请求已收到，订单正在处理中
	return nil
}
