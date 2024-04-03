package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResult struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Result  interface{} `json:"result"`
}

func NewJsonResult(code int, success bool, msg string, data interface{}) *JsonResult {
	return &JsonResult{
		Code:    code,
		Success: success,
		Msg:     msg,
		Result:  data,
	}
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(
		http.StatusOK,
		NewJsonResult(http.StatusOK, true, msg, data),
	)
}

func SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		NewJsonResult(http.StatusOK, true, "success", data),
	)
}

func Failure(c *gin.Context, msg string, data interface{}) {
	c.JSON(
		http.StatusInternalServerError,
		NewJsonResult(
			http.StatusInternalServerError, false, msg, data),
	)
}

func FailureWithCode(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(
		code,
		NewJsonResult(
			code, false, msg, data),
	)
}

func FailureWithData(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusInternalServerError,
		NewJsonResult(
			http.StatusInternalServerError, false, "failure", data),
	)
}
