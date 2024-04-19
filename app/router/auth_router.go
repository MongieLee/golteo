package router

import (
	"ginl/app/controller"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(e *gin.Engine) {
	authController := &controller.AuthController{}
	authGroup := e.Group("/api/v1/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/refreshToken", authController.RefreshToken)
	}
}
