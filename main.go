package main

import (
	"ginl/app/router"
	"ginl/config"
	"ginl/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.InitFlag()
	config.InitViperConfig()
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
