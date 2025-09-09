// file: backend/api/seckill.go
package api

import (
	"net/http"
	"seckill-project/backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SeckillHandler 处理秒杀请求
func SeckillHandler(c *gin.Context) {
	// 1. 从 JWT 中间件获取用户 ID
	// 我们在这里使用 c.MustGet，如果 userID 不存在，程序会 panic
	// 因为 JWTMiddleware 应该确保 userID 总是存在的
	userID := c.MustGet("userID").(uint)

	// 2. 从 URL 中获取活动 ID
	activityIDStr := c.Param("id")
	activityID, err := strconv.ParseUint(activityIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的活动ID"})
		return
	}

	// 3. 调用 Service 层的高并发版秒杀逻辑
	err = service.ExecuteSeckillHighConcurrency(userID, uint(activityID))
	if err != nil {
		// 根据业务错误返回相应的提示
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// 4. 秒杀请求已成功接收
	c.JSON(http.StatusOK, gin.H{
		"message": "抢购请求已收到，订单正在处理中！",
	})
}
