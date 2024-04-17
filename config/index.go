package config

import (
	"ginl/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RedisDb struct {
	Client *redis.Client
	PubSub []*redis.PubSub
}

var Db *gorm.DB
var Rdb RedisDb

func Init() {
	InitFlag()
	InitViperConfig()
	utils.InitTrans("zh")
	InitDb()
	InitRate()
	InitRedisDb()
	if !CustomConfig.AppDebug {
		gin.SetMode(gin.ReleaseMode)
	}
}
