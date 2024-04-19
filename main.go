package main

import (
	"ginl/app/router"
	"ginl/config"
	"github.com/gin-gonic/gin"
)

func main() {
	defer config.CloseAll()
	config.Init()
	r := gin.Default()
	router.InitRouters(r)
	err := r.Run()
	if err != nil {
		return
	}
}
