package router

import (
	"context"
	"ginl/app/controller"
	"ginl/config"
	"ginl/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"
)

func InitSocketRouter(e *gin.Engine) {
	g := e.Group("/api/v1")
	socketController := controller.SocketController{}
	g.GET("/getOnlineUsersCount", socketController.GetOnlineUserCount)
	g.GET("/socket.io", socketController.InitWebSocket)
	inspectRedisSocketConnValid()
}

func inspectRedisSocketConnValid() {
	cmd := config.Rdb.Client.ZRangeArgsWithScores(context.Background(), redis.ZRangeArgs{
		Key:   "online_uids",
		Start: "0",
		Stop:  time.Now().Unix(),
	})
	if cmd.Err() != nil {
		utils.ErrorF("检查redis链接集合失败")
	}
	if len(cmd.Val()) > 0 {
		cmd := config.Rdb.Client.Del(context.Background(), "online_uids")
		if cmd.Err() == nil {
			utils.InfoF("redis的online_uids被删除了")
		}
	}
}
