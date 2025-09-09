package dao

import (
	"context"
	"seckill-project/backend/config"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.DB,
	})
	_, err = RDB.Ping(Ctx).Result()
	return err
}
