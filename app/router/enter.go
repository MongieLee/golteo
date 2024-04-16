package router

import (
	"ginl/config"
	"ginl/middleware"
	"ginl/service/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initMiddleware(r *gin.Engine) {
	r.Use(middleware.CorsHandler)
	if config.CustomConfig.RateConfig.Enable {
		r.Use(middleware.RateHandler)
	}
}

func initStaticResource(e *gin.Engine) {
	e.Static("/files", "./uploads")
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
	InitAuthRoute(e.Group("/api/v1"))
}
