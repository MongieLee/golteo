package config

import (
	"context"
	"errors"
	"fmt"
	"ginl/common"
	"github.com/redis/go-redis/v9"
	"log"
	"reflect"
	"strings"
	"time"
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

func (r RedisDb) Set(key string, data interface{}, expireTime int64) error {
	dataStr := ""
	if reflect.TypeOf(data).Kind() == reflect.String {
		dataStr = data.(string)
	} else if reflect.TypeOf(data).Kind() == reflect.Ptr && reflect.TypeOf(data).Elem().Kind() == reflect.String {
		dataStr = *data.(*string)
	} else {
		dataStr = common.JsonEncode(data)
	}
	if dataStr == "" {
		msg := fmt.Sprintf("尝试写入Key:%s的数据为空字符串，不予写入,", key)
		return errors.New(msg)
	}
	err := r.Client.Set(context.Background(), strings.ToLower(key), dataStr, time.Duration(expireTime)*time.Second).Err()
	return err
}

func (r RedisDb) Get(key string, val interface{}) error {
	redisData, err := r.Client.Get(context.Background(), strings.ToLower(key)).Result()
	if err != nil {
		return err
	}
	if reflect.TypeOf(val).Kind() == reflect.String {
		val = redisData
		return err
	}

	if reflect.TypeOf(val).Kind() == reflect.Ptr && reflect.TypeOf(val).Elem().Kind() == reflect.String {
		*val.(*string) = redisData
		return err
	}
	return common.JsonDecode(redisData, &val)
}
