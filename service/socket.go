package service

import (
	"context"
	"fmt"
	"ginl/config"
	"ginl/utils"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var Upgrade = websocket.Upgrader{
	// 校验请求来源，返回true直接跳过
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 心跳包发送时间
var pingTime = 5 * time.Second

var redisChannelName = fmt.Sprintf("%s_%d", "gin-pure-listenter", config.CustomConfig.Redis.Database)

func InitSocket(conn *websocket.Conn) {
	in := make(chan []byte)
	pingTicker := time.NewTicker(pingTime)
	stop := make(chan struct{})

	go func() {
		// redis订阅监听
		pubSub := config.Rdb.Client.Subscribe(context.Background(), redisChannelName)
		// 等待订阅确认
		_, err := pubSub.Receive(context.Background())
		if err != nil {
			panic(err)
		}
		ch := pubSub.Channel()
		for msg := range ch {
			//var pushData map[string]any
			utils.InfoF("监听redis收到信息msg.Payload：")
			utils.InfoF(msg.Payload)
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			if err != nil {
				utils.ErrorF("写入失败，%v", err)
			}
			// 拿到数据后，再筛分是否要进行数据推送
		}
	}()

	go func() {
		for {
			// ReadMessage会阻塞
			_, msg, err := conn.ReadMessage()
			if err != nil {
				// 客户端主动断开时，也会收到错误，websocket: close 1005 (no status)
				utils.ErrorF("[ReadMessage]方法读取时发生异常:%v", err)
				close(stop)
				break
			}
			in <- msg
		}
	}()

	for {
		select {
		case <-pingTicker.C:
			//err := conn.WriteMessage(websocket.TextMessage, []byte("{\"event\":\"gin pro\"}"))
			//if err != nil {
			//	log.Printf("写入失败，%v", err)
			//}
			//log.Printf("心跳包发送成功")
			// 这是主业务逻辑服务器的某段业务代码，更新redis缓存后，再发布消息让推送服务能监听到
			config.Rdb.Client.Publish(context.Background(), redisChannelName,
				"{\"event\":\"someEvent\",\"data\":\"[]\"}")
		case message := <-in:
			{
				utils.InfoF("接收到数据：%v\n", string(message))
			}
		case <-stop:
			// stop管道有值时，说明出现异常需要停止
			return
		}
	}

}
