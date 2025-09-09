package api

import (
	"net/http"
	"seckill-project/backend/internal/model"
	"seckill-project/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// CreateProductHandler 创建商品
func CreateProductHandler(c *gin.Context) {
	var p model.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数无效"})
		return
	}

	if err := service.CreateProduct(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "商品创建成功",
		"product": p,
	})
}
