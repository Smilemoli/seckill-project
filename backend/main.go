// file: backend/main.go
package main

import (
	"fmt"
	"log"
	"seckill-project/backend/api"
	"seckill-project/backend/config"
	"seckill-project/backend/internal/dao"
	"seckill-project/backend/internal/middleware"
	"seckill-project/backend/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 初始化配置

	config.Init()

	// 2. 初始化数据库连接
	if err := dao.InitMySQL(); err != nil {
		log.Fatalf("dao.InitMySQL() failed, err: %v", err)
	}
	fmt.Println("MySQL connected successfully.")

	if err := dao.InitRedis(); err != nil {
		log.Fatalf("dao.InitRedis() failed, err: %v", err)
	}
	fmt.Println("Redis connected successfully.")

	go service.StartOrderConsumer()
	fmt.Println("Order consumer started.")

	// 3. 初始化 Gin 引擎
	r := gin.Default()

	// 创建一个 API 版本 v1 的路由组
	v1 := r.Group("/api/v1")
	{
		// 用户注册和登录
		v1.POST("/register", api.RegisterHandler)
		v1.POST("/login", api.LoginHandler)

		// 秒杀活动列表是公开的，任何人都可以看
		v1.GET("/activities", api.GetActivitiesHandler)

		// 需要认证的路由组
		authed := v1.Group("")
		authed.Use(middleware.JWTMiddleware())
		{
			// 用户信息
			authed.GET("/profile", api.ProfileHandler)

			// 商品和活动管理 (简化为普通用户即可操作)
			authed.POST("/products", api.CreateProductHandler)
			authed.POST("/activities", api.CreateActivityHandler)
			authed.POST("/seckill/:id", api.SeckillHandler)
		}
	}

	// 5. 启动服务
	fmt.Println("Server is running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server run failed, err: %v", err)
	}
}
