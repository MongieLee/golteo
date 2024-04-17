package middleware

import (
	"ginl/service/result"
	"github.com/gin-gonic/gin"
	"log"
)

func ExceptionMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("catch Exception.%v", err)
			result.Failure(c, "Internal Server Error", gin.H{})
			c.Abort()
		}
	}()
	c.Next()
}
