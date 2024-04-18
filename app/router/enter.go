package router

import (
	"ginl/app/middleware"
	"ginl/config"
	"ginl/service"
	"ginl/service/result"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func initMiddleware(r *gin.Engine) {
	r.Use(middleware.CorsHandler)
	if config.CustomConfig.RateConfig.Enable {
		r.Use(middleware.RateHandler)
	}
	r.Use(middleware.ExceptionMiddleware)
}

func initStaticResource(e *gin.Engine) {
	e.Static("/files", "./uploads")
}

func InitWebsocketServer(e *gin.Engine) {
	e.GET("/ws", func(c *gin.Context) {
		conn, err := service.Upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("websocket创建失败，%v", err)
			return
		}
		service.InitSocket(conn)
	})
}

func InitRouters(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		result.FailureWithCode(c, http.StatusNotFound, "资源不存在", gin.H{})
	})
	initMiddleware(e)
	initStaticResource(e)
	group := e.Group("/api/v1")
	group.Use(middleware.AuthHandler)
	{
		InitUserRoute(group)
		InitFileRoute(group)
	}
	InitAuthRoute(group)
	InitWebsocketServer(e)
}
