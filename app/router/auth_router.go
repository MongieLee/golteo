package router

import (
	"ginl/app/controller"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(e *gin.RouterGroup) {
	authController := &controller.AuthController{}
	authGroup := e.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/refreshToken", authController.RefreshToken)
	}
}
