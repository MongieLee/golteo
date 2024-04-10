package router

import (
	"ginl/app/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(e *gin.RouterGroup) {
	userController := &controller.UserController{}
	{
		e.GET("/users/:id", userController.GetUserById)
		e.GET("/users", userController.GetUsers)
		e.PATCH("/users", userController.ModifyUser)
		e.DELETE("/users/:id", userController.DeleteUser)
	}
}
