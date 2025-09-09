package middleware

import (
	"errors"
	"net/http"
	"seckill-project/backend/config"
	"seckill-project/backend/internal/types"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTMiddleware 创建一个 JWT 认证中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求头中获取 Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "请求未携带Token"})
			return
		}

		// 2. 校验 Token 格式，必须是 "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token格式错误"})
			return
		}

		tokenString := parts[1]

		// 3. 解析和校验 Token
		claims := &types.CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Conf.JWT.Secret), nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token已过期"})
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的Token"})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的Token"})
			return
		}

		// 4. Token 校验通过，将用户信息存入 Context
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		// 5. 放行请求
		c.Next()
	}
}
