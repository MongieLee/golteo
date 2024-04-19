package controller

import (
	"context"
	"ginl/config"
	"ginl/service"
	"ginl/service/result"
	"ginl/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type SocketController struct {
}

func (skc SocketController) GetOnlineUserCount(c *gin.Context) {
	ct := time.Now().Unix()
	cmd := config.Rdb.Client.ZCount(context.Background(), "online_uids", "0", strconv.Itoa(int(ct)))
	if err := cmd.Err(); err != nil {
		result.FailureWithData(c, gin.H{})
	}
	i, err := cmd.Result()
	if err != nil {
		result.FailureWithData(c, gin.H{})
	}
	result.SuccessWithData(c, gin.H{
		"onlineCount": i,
	})
}

func (skc SocketController) InitWebSocket(c *gin.Context) {
	conn, err := service.Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		utils.ErrorF("websocket创建失败，%v", err)
		return
	}
	service.InitSocket(conn)
}
