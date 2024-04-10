package router

import (
	"ginl/middleware"
	"github.com/gin-gonic/gin"
)

func initMiddleware(r *gin.Engine) {
	r.Use(middleware.CorsHandler)
}

func initStaticResource(e *gin.Engine) {
	e.Static("/files", "./uploads")
}

func InitRouters(e *gin.Engine) {
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
