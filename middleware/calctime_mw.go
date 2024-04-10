package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func CalcTimeHandler(c *gin.Context) {
	start := time.Now()
	c.Next()
	log.Print("此次请求花费时间：" + fmt.Sprintf("%v", time.Since(start)))
}
