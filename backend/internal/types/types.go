package types

import "github.com/golang-jwt/jwt/v5"

// CustomClaims 自定义声明类型
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
