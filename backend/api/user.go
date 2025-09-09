// file: backend/api/user.go
package api

import (
	"net/http"
	"seckill-project/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// RegisterRequest 定义了注册请求的 JSON 结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest 定义了登录请求的 JSON 结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterHandler 处理用户注册请求
func RegisterHandler(c *gin.Context) {
	var req RegisterRequest
	// 1. 绑定并校验 JSON 请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数无效"})
		return
	}

	// 2. 调用 service 层的注册逻辑
	if err := service.Register(req.Username, req.Password); err != nil {
		// 注册过程中出现错误（比如用户已存在）
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. 注册成功，返回响应
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// LoginHandler 处理用户登录请求
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数无效"})
		return
	}

	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}

// ProfileHandler 获取用户信息
func ProfileHandler(c *gin.Context) {
	// 从中间件设置的 Context 中获取用户信息
	userID, _ := c.Get("userID")
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{
		"message":  "认证成功",
		"userID":   userID,
		"username": username,
	})
}
