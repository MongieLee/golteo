package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"log"
)

type bodyLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWrite) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinBodyLogHandler(c *gin.Context) {
	blw := &bodyLogWrite{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	log.Print("response body:" + blw.body.String())
}
