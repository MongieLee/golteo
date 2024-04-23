package ginl

import (
	"ginl/app/router"
	"ginl/config"
	"github.com/gin-gonic/gin"
)

func NewServer() {
	defer config.CloseAll()
	config.Init()
	r := gin.Default()
	router.InitRouters(r)
	err := r.Run()
	if err != nil {
		return
	}
}
