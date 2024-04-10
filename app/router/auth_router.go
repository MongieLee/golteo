package router

import (
	"ginl/app/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func InitAuthRoute(e *gin.RouterGroup) {
	authController := &controller.AuthController{}
	authGroup := e.Group("/auth", func(c *gin.Context) {
		log.Println("auth组才有的中间件")
		c.Next()
	})
	authGroup.Use(func(c *gin.Context) {
		log.Println("auth组才有的中间件222")
		// 直接使用c是并发不安全的，通过c.Copy获取一份上下午文的副本
		c.Set("fuck", "aaaaafuck")
		contextCopy := c.Copy()
		c.Set("fuck", "12312312")
		go func() {
			log.Println("在goroutine中开启的输出打印" + contextCopy.Request.URL.Path)
			value, exists := contextCopy.Get("fuck")
			if !exists {
				return
			}
			s, ok := value.(string)
			if !ok {
				return
			}
			log.Println("在goroutine中开启的输出打印" + s)
		}()
		c.Next()
	})
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/refreshToken", authController.RefreshToken)
	}
}
