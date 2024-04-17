package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		CustomConfig.Mysql.User,
		CustomConfig.Mysql.Password,
		CustomConfig.Mysql.Hostname,
		CustomConfig.Mysql.Port,
		CustomConfig.Mysql.Database)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("链接Mysql出错，错误信息：%v", err)
	}
}
