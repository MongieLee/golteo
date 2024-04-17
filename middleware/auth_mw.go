package middleware

import (
	"ginl/service/result"
	"ginl/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthHandler(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		result.FailureWithCode(c, http.StatusBadRequest, "鉴权失败，token不能为空", gin.H{})
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		result.FailureWithCode(c, http.StatusBadRequest, "请求头中的auth格式不对", gin.H{})
		return
	}
	customClaims, err := utils.ParseJWTToken(parts[1])
	if err != nil {
		result.FailureWithCode(c, http.StatusBadRequest, "token校验失败", gin.H{})
		c.Abort()
		return
	}
	c.Set("userInfo", customClaims)
	c.Next()
}
