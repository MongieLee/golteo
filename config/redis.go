package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func InitRedisDb() {
	Rdb.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", CustomConfig.Redis.Hostname, CustomConfig.Redis.Port),
		Password: CustomConfig.Redis.Password,
		DB:       CustomConfig.Redis.Database,
		PoolSize: 20,
	})
	_, err := Rdb.Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis connect failed,err:%v", err)
	}
}
