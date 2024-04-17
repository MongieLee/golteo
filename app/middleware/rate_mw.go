package middleware

import (
	"ginl/config"
	"ginl/service/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RateHandler(c *gin.Context) {
	if !config.Limiter.Allow() {
		result.FailureWithCode(c, http.StatusTooManyRequests, "try again later", gin.H{})
		c.Abort()
	}
	c.Next()
}
