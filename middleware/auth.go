package middleware

import (
	"ginl/service/result"
	"ginl/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthHandler(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		result.FailureWithData(c, gin.H{
			"code": 2003,
			"msg":  "token为空",
		})
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		result.FailureWithData(c, gin.H{
			"code": 20004,
			"msg":  "请求头中的auth格式不对",
		})
		return
	}
	customClaims, err := utils.ParseJWTToken(parts[1])
	if err != nil {
		result.FailureWithData(c, gin.H{
			"code": 20005,
			"msg":  "无效token",
		})
		c.Abort()
		return
	}
	c.Set("userInfo", customClaims)
	c.Next()
}
