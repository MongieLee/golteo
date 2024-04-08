package db

import (
	"database/sql"
	"fmt"
	"ginl/config"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Db *sql.DB
var GormDb *gorm.DB

func InitDb() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", config.CustomConfig.Mysql.User, config.CustomConfig.Mysql.Password, config.CustomConfig.Mysql.Hostname, config.CustomConfig.Mysql.Port, config.CustomConfig.Mysql.Database)
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	GormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 格式化时间，相当于yyyy-MM-dd HH:mm:ss
const localTimeFormat string = "2006-01-02 15:04:05"

// LocalTime 自定义一个类型，本质上是time.Time，但是重写该方法的MarshalJSON方法来改变返回值
type LocalTime time.Time

// MarshalJSON 自定义成序列化JSON内容
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	t2 := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", t2.Format(localTimeFormat))), nil
}
