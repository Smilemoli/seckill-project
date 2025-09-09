package api

import (
	"net/http"
	"seckill-project/backend/internal/model"
	"seckill-project/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// CreateActivityHandler 创建秒杀活动
func CreateActivityHandler(c *gin.Context) {
	var activity model.SeckillActivity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusAlreadyReported, gin.H{
			"error": "参数无效",
		})
		return
	}
	if err := service.CreateActivity(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "秒杀活动创建成功",
		"activity": activity,
	})
}

// GetActivitiesHandler 获取秒杀活动列表
func GetActivitiesHandler(c *gin.Context) {
	activities, err := service.GetActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"activities": activities,
	})
}
