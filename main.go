package main

import (
	"ginl/app/router"
	"ginl/config"
	"ginl/db"
	"ginl/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.InitFlag()
	config.InitViperConfig()
	transErr := utils.InitTrans("zh")
	if transErr != nil {
		log.Printf("init trans failed,err : %v\n", transErr)
	}
	dbErr := db.InitDb()
	if dbErr != nil {
		log.Println(dbErr.Error())
		return
	}
	r := gin.Default()
	router.InitRouters(r)
	err := r.Run()
	if err != nil {
		return
	}
}
